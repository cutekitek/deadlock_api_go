#!/usr/bin/env python3
import re
import sys

def generate_code(proto_content):
    # Search for the enum block.
    enum_pattern = re.compile(r"enum\s+EGCCitadelClientMessages\s*{(.*?)}", re.DOTALL)
    enum_match = enum_pattern.search(proto_content)
    if not enum_match:
        sys.exit("Error: Enum EGCCitadelClientMessages not found in proto file.")
    enum_body = enum_match.group(1)

    # Find all entries matching k_EMsgClientToGCXXX (XXX can be any word)
    entry_pattern = re.compile(r"(k_EMsgClientToGC(\w+))")
    entries = entry_pattern.findall(enum_body)
    if not entries:
        sys.exit("Error: No entries found matching pattern 'k_EMsgClientToGCXXX'.")

    command_functions = []
    response_structs = []
    response_cases = []

    for full_name, suffix in entries:
        # Check if the entry is a response (ends with "Response")
        if suffix.endswith("Response"):
            # Derive the base name (without the "Response" suffix) for the struct.
            base_name = suffix[:-len("Response")]
            # Generate the struct definition.
            struct_def = f"""type {base_name}Response struct {{
    *deadlock.CMsgClientToGC{suffix}
}}
"""
            response_structs.append(struct_def)

            # Generate the case for the ParseGCResponse function.
            case_code = f"""case deadlock.EGCCitadelClientMessages_k_EMsgClientToGC{suffix}:
    msg := &deadlock.CMsgClientToGC{suffix}{{}}
    err := proto.Unmarshal(packet.Body, msg)
    if err != nil {{
        return nil, err
    }}
    return {base_name}Response{{msg}}, nil"""
            response_cases.append(case_code)
        else:
            # Generate a normal function for non-response messages.
            func_template = f"""func (h *Handler) {suffix}(msg *deadlock.CMsgClientToGC{suffix}) {{
    h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGC{suffix}), msg))
}}
"""
            command_functions.append(func_template)

    commands_output = "\n".join(command_functions)
    responses_structs_output = "\n".join(response_structs)
    response_cases_output = "\n\n".join(response_cases)

    parse_function = f"""func (h *handler) ParseGCResponse(packet *gamecoordinator.GCPacket) (interface{{}}, error) {{
    switch deadlock.EGCCitadelClientMessages(packet.MsgType) {{
{response_cases_output}
    }}
    return nil, fmt.Errorf("unknown message type: %d", packet.MsgType)
}}
"""

    return commands_output, responses_structs_output, parse_function

def main():
    if len(sys.argv) < 2:
        sys.exit("Usage: python codegen.py <proto_file.proto>")
    proto_filename = sys.argv[1]

    try:
        with open(proto_filename, "r") as f:
            proto_content = f.read()
    except Exception as e:
        sys.exit(f"Error reading file: {e}")

    commands_output, responses_structs_output, parse_function_output = generate_code(proto_content)

    # Output the generated code.
    print("// Generated command functions:")
    print(commands_output)
    print("\n// Generated response structs:")
    print(responses_structs_output)
    print("\n// Generated ParseGCResponse function:")
    print(parse_function_output)

if __name__ == "__main__":
    main()
