package api

import (
	"deadlock_analyzer/proto/gc/deadlock"
	"fmt"

	"github.com/0xAozora/go-steam/protocol/gamecoordinator"
	"google.golang.org/protobuf/proto"
)

func (h *Client) StartMatchmaking(msg *deadlock.CMsgClientToGCStartMatchmaking) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmaking), msg))
}

func (h *Client) StopMatchmaking(msg *deadlock.CMsgClientToGCStopMatchmaking) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmaking), msg))
}

func (h *Client) LeaveLobby(msg *deadlock.CMsgClientToGCLeaveLobby) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobby), msg))
}

func (h *Client) IsInMatchmaking(msg *deadlock.CMsgClientToGCIsInMatchmaking) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmaking), msg))
}

func (h *Client) DevSetMMBias(msg *deadlock.CMsgClientToGCDevSetMMBias) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCDevSetMMBias), msg))
}

func (h *Client) GetProfileCard(msg *deadlock.CMsgClientToGCGetProfileCard) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard), msg))
}

func (h *Client) UpdateRoster(msg *deadlock.CMsgClientToGCUpdateRoster) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRoster), msg))
}

func (h *Client) ModifyDevAnnouncements(msg *deadlock.CMsgClientToGCModifyDevAnnouncements) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCModifyDevAnnouncements), msg))
}

func (h *Client) ReplacementSDRTicket(msg *deadlock.CMsgClientToGCReplacementSDRTicket) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicket), msg))
}

func (h *Client) SetServerConVar(msg *deadlock.CMsgClientToGCSetServerConVar) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVar), msg))
}

func (h *Client) SpectateLobby(msg *deadlock.CMsgClientToGCSpectateLobby) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobby), msg))
}

func (h *Client) GetMatchHistory(msg *deadlock.CMsgClientToGCGetMatchHistory) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistory), msg))
}

func (h *Client) SpectateUser(msg *deadlock.CMsgClientToGCSpectateUser) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUser), msg))
}

func (h *Client) PartyCreate(msg *deadlock.CMsgClientToGCPartyCreate) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreate), msg))
}

func (h *Client) PartyLeave(msg *deadlock.CMsgClientToGCPartyLeave) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeave), msg))
}

func (h *Client) PartyJoin(msg *deadlock.CMsgClientToGCPartyJoin) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoin), msg))
}

func (h *Client) PartyAction(msg *deadlock.CMsgClientToGCPartyAction) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyAction), msg))
}

func (h *Client) PartyStartMatch(msg *deadlock.CMsgClientToGCPartyStartMatch) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatch), msg))
}

func (h *Client) PartyInviteUser(msg *deadlock.CMsgClientToGCPartyInviteUser) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUser), msg))
}

func (h *Client) PartyJoinViaCode(msg *deadlock.CMsgClientToGCPartyJoinViaCode) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCode), msg))
}

func (h *Client) PartySetReadyState(msg *deadlock.CMsgClientToGCPartySetReadyState) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyState), msg))
}

func (h *Client) GetAccountStats(msg *deadlock.CMsgClientToGCGetAccountStats) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStats), msg))
}

func (h *Client) GetMatchMetaData(msg *deadlock.CMsgClientToGCGetMatchMetaData) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaData), msg))
}

func (h *Client) DevAction(msg *deadlock.CMsgClientToGCDevAction) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCDevAction), msg))
}

func (h *Client) RecordClientEvents(msg *deadlock.CMsgClientToGCRecordClientEvents) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEvents), msg))
}

func (h *Client) SetNewPlayerProgress(msg *deadlock.CMsgClientToGCSetNewPlayerProgress) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgress), msg))
}

func (h *Client) UpdateAccountSync(msg *deadlock.CMsgClientToGCUpdateAccountSync) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSync), msg))
}

func (h *Client) GetHeroChoice(msg *deadlock.CMsgClientToGCGetHeroChoice) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoice), msg))
}

func (h *Client) UnlockHero(msg *deadlock.CMsgClientToGCUnlockHero) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHero), msg))
}

func (h *Client) BookUnlock(msg *deadlock.CMsgClientToGCBookUnlock) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlock), msg))
}

func (h *Client) GetBook(msg *deadlock.CMsgClientToGCGetBook) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetBook), msg))
}

func (h *Client) SubmitPlaytestUser(msg *deadlock.CMsgClientToGCSubmitPlaytestUser) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUser), msg))
}

func (h *Client) UpdateHeroBuild(msg *deadlock.CMsgClientToGCUpdateHeroBuild) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuild), msg))
}

func (h *Client) FindHeroBuilds(msg *deadlock.CMsgClientToGCFindHeroBuilds) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuilds), msg))
}

func (h *Client) ReportPlayerFromMatch(msg *deadlock.CMsgClientToGCReportPlayerFromMatch) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatch), msg))
}

func (h *Client) GetAccountMatchReports(msg *deadlock.CMsgClientToGCGetAccountMatchReports) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReports), msg))
}

func (h *Client) DeleteHeroBuild(msg *deadlock.CMsgClientToGCDeleteHeroBuild) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuild), msg))
}

func (h *Client) GetActiveMatches(msg *deadlock.CMsgClientToGCGetActiveMatches) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatches), msg))
}

func (h *Client) GetDiscordLink(msg *deadlock.CMsgClientToGCGetDiscordLink) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLink), msg))
}

func (h *Client) PartySetMode(msg *deadlock.CMsgClientToGCPartySetMode) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartySetMode), msg))
}

func (h *Client) GrantForumAccess(msg *deadlock.CMsgClientToGCGrantForumAccess) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccess), msg))
}

func (h *Client) ModeratorRequest(msg *deadlock.CMsgClientToGCModeratorRequest) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequest), msg))
}

func (h *Client) GetFriendGameStatus(msg *deadlock.CMsgClientToGCGetFriendGameStatus) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatus), msg))
}

func (h *Client) UpdateHeroBuildPreference(msg *deadlock.CMsgClientToGCUpdateHeroBuildPreference) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreference), msg))
}

func (h *Client) GetOldHeroBuildData(msg *deadlock.CMsgClientToGCGetOldHeroBuildData) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetOldHeroBuildData), msg))
}

func (h *Client) UpdateSpectatorStatus(msg *deadlock.CMsgClientToGCUpdateSpectatorStatus) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateSpectatorStatus), msg))
}

func (h *Client) CommendPlayerFromMatch(msg *deadlock.CMsgClientToGCCommendPlayerFromMatch) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatch), msg))
}

func (h *Client) GetHeroMMRRankings(msg *deadlock.CMsgClientToGCGetHeroMMRRankings) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankings), msg))
}

func (h *Client) GetLeaderboard(msg *deadlock.CMsgClientToGCGetLeaderboard) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboard), msg))
}

func (h *Client) GetAccountLeaderboards(msg *deadlock.CMsgClientToGCGetAccountLeaderboards) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboards), msg))
}

func (h *Client) SetHideHolidayModels(msg *deadlock.CMsgClientToGCSetHideHolidayModels) {
	h.Client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppId, uint32(deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetHideHolidayModels), msg))
}

// Generated response structs:
type StartMatchmakingResponse struct {
	*deadlock.CMsgClientToGCStartMatchmakingResponse
}

type StopMatchmakingResponse struct {
	*deadlock.CMsgClientToGCStopMatchmakingResponse
}

type LeaveLobbyResponse struct {
	*deadlock.CMsgClientToGCLeaveLobbyResponse
}

type IsInMatchmakingResponse struct {
	*deadlock.CMsgClientToGCIsInMatchmakingResponse
}

type GetProfileCardResponse struct {
	*deadlock.CMsgClientToGCGetProfileCard
}

type UpdateRosterResponse struct {
	*deadlock.CMsgClientToGCUpdateRosterResponse
}

type ModifyDevAnnouncementsResponse struct {
	*deadlock.CMsgClientToGCModifyDevAnnouncementsResponse
}

type ReplacementSDRTicketResponse struct {
	*deadlock.CMsgClientToGCReplacementSDRTicketResponse
}

type SetServerConVarResponse struct {
	*deadlock.CMsgClientToGCSetServerConVarResponse
}

type SpectateLobbyResponse struct {
	*deadlock.CMsgClientToGCSpectateLobbyResponse
}

type PostMatchSurveyResponse struct {
	*deadlock.CMsgClientToGCPostMatchSurveyResponse
}

type GetMatchHistoryResponse struct {
	*deadlock.CMsgClientToGCGetMatchHistoryResponse
}

type SpectateUserResponse struct {
	*deadlock.CMsgClientToGCSpectateUserResponse
}

type PartyCreateResponse struct {
	*deadlock.CMsgClientToGCPartyCreateResponse
}

type PartyLeaveResponse struct {
	*deadlock.CMsgClientToGCPartyLeaveResponse
}

type PartyJoinResponse struct {
	*deadlock.CMsgClientToGCPartyJoinResponse
}

type PartyActionResponse struct {
	*deadlock.CMsgClientToGCPartyActionResponse
}

type PartyStartMatchResponse struct {
	*deadlock.CMsgClientToGCPartyStartMatchResponse
}

type PartyInviteUserResponse struct {
	*deadlock.CMsgClientToGCPartyInviteUserResponse
}

type PartyJoinViaCodeResponse struct {
	*deadlock.CMsgClientToGCPartyJoinViaCodeResponse
}

type PartySetReadyStateResponse struct {
	*deadlock.CMsgClientToGCPartySetReadyStateResponse
}

type GetAccountStatsResponse struct {
	*deadlock.CMsgClientToGCGetAccountStatsResponse
}

type GetMatchMetaDataResponse struct {
	*deadlock.CMsgClientToGCGetMatchMetaDataResponse
}

type DevActionResponse struct {
	*deadlock.CMsgClientToGCDevActionResponse
}

type RecordClientEventsResponse struct {
	*deadlock.CMsgClientToGCRecordClientEventsResponse
}

type SetNewPlayerProgressResponse struct {
	*deadlock.CMsgClientToGCSetNewPlayerProgressResponse
}

type UpdateAccountSyncResponse struct {
	*deadlock.CMsgClientToGCUpdateAccountSyncResponse
}

type GetHeroChoiceResponse struct {
	*deadlock.CMsgClientToGCGetHeroChoiceResponse
}

type UnlockHeroResponse struct {
	*deadlock.CMsgClientToGCUnlockHeroResponse
}

type BookUnlockResponse struct {
	*deadlock.CMsgClientToGCBookUnlockResponse
}

type GetBookResponse struct {
	*deadlock.CMsgClientToGCGetBookResponse
}

type SubmitPlaytestUserResponse struct {
	*deadlock.CMsgClientToGCSubmitPlaytestUserResponse
}

type UpdateHeroBuildResponse struct {
	*deadlock.CMsgClientToGCUpdateHeroBuildResponse
}

type FindHeroBuildsResponse struct {
	*deadlock.CMsgClientToGCFindHeroBuildsResponse
}

type ReportPlayerFromMatchResponse struct {
	*deadlock.CMsgClientToGCReportPlayerFromMatchResponse
}

type GetAccountMatchReportsResponse struct {
	*deadlock.CMsgClientToGCGetAccountMatchReportsResponse
}

type DeleteHeroBuildResponse struct {
	*deadlock.CMsgClientToGCDeleteHeroBuildResponse
}

type GetActiveMatchesResponse struct {
	*deadlock.CMsgClientToGCGetActiveMatchesResponse
}

type GetDiscordLinkResponse struct {
	*deadlock.CMsgClientToGCGetDiscordLinkResponse
}

type PartySetModeResponse struct {
	*deadlock.CMsgClientToGCPartySetModeResponse
}

type GrantForumAccessResponse struct {
	*deadlock.CMsgClientToGCGrantForumAccessResponse
}

type ModeratorRequestResponse struct {
	*deadlock.CMsgClientToGCModeratorRequestResponse
}

type GetFriendGameStatusResponse struct {
	*deadlock.CMsgClientToGCGetFriendGameStatusResponse
}

type UpdateHeroBuildPreferenceResponse struct {
	*deadlock.CMsgClientToGCUpdateHeroBuildPreferenceResponse
}

type GetOldHeroBuildDataResponse struct {
	*deadlock.CMsgClientToGCGetOldHeroBuildDataResponse
}

type CommendPlayerFromMatchResponse struct {
	*deadlock.CMsgClientToGCCommendPlayerFromMatchResponse
}

type GetHeroMMRRankingsResponse struct {
	*deadlock.CMsgClientToGCGetHeroMMRRankingsResponse
}

type GetLeaderboardResponse struct {
	*deadlock.CMsgClientToGCGetLeaderboardResponse
}

type GetAccountLeaderboardsResponse struct {
	*deadlock.CMsgClientToGCGetAccountLeaderboardsResponse
}

type SetHideHolidayModelsResponse struct {
	*deadlock.CMsgClientToGCSetHideHolidayModelsResponse
}

// Generated ParseGCResponse function:
func (h *Client) ParseGCResponse(packet *gamecoordinator.GCPacket) (interface{}, error) {
	switch deadlock.EGCCitadelClientMessages(packet.MsgType) {
	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmakingResponse:
		msg := &deadlock.CMsgClientToGCStartMatchmakingResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return StartMatchmakingResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmakingResponse:
		msg := &deadlock.CMsgClientToGCStopMatchmakingResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return StopMatchmakingResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobbyResponse:
		msg := &deadlock.CMsgClientToGCLeaveLobbyResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return LeaveLobbyResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmakingResponse:
		msg := &deadlock.CMsgClientToGCIsInMatchmakingResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return IsInMatchmakingResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCardResponse:
		msg := &deadlock.CMsgClientToGCGetProfileCard{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetProfileCardResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRosterResponse:
		msg := &deadlock.CMsgClientToGCUpdateRosterResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return UpdateRosterResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCModifyDevAnnouncementsResponse:
		msg := &deadlock.CMsgClientToGCModifyDevAnnouncementsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return ModifyDevAnnouncementsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicketResponse:
		msg := &deadlock.CMsgClientToGCReplacementSDRTicketResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return ReplacementSDRTicketResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVarResponse:
		msg := &deadlock.CMsgClientToGCSetServerConVarResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SetServerConVarResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobbyResponse:
		msg := &deadlock.CMsgClientToGCSpectateLobbyResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SpectateLobbyResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPostMatchSurveyResponse:
		msg := &deadlock.CMsgClientToGCPostMatchSurveyResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PostMatchSurveyResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistoryResponse:
		msg := &deadlock.CMsgClientToGCGetMatchHistoryResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetMatchHistoryResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUserResponse:
		msg := &deadlock.CMsgClientToGCSpectateUserResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SpectateUserResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreateResponse:
		msg := &deadlock.CMsgClientToGCPartyCreateResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyCreateResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeaveResponse:
		msg := &deadlock.CMsgClientToGCPartyLeaveResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyLeaveResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinResponse:
		msg := &deadlock.CMsgClientToGCPartyJoinResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyJoinResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyActionResponse:
		msg := &deadlock.CMsgClientToGCPartyActionResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyActionResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatchResponse:
		msg := &deadlock.CMsgClientToGCPartyStartMatchResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyStartMatchResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUserResponse:
		msg := &deadlock.CMsgClientToGCPartyInviteUserResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyInviteUserResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCodeResponse:
		msg := &deadlock.CMsgClientToGCPartyJoinViaCodeResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartyJoinViaCodeResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyStateResponse:
		msg := &deadlock.CMsgClientToGCPartySetReadyStateResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartySetReadyStateResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStatsResponse:
		msg := &deadlock.CMsgClientToGCGetAccountStatsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetAccountStatsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaDataResponse:
		msg := &deadlock.CMsgClientToGCGetMatchMetaDataResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetMatchMetaDataResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCDevActionResponse:
		msg := &deadlock.CMsgClientToGCDevActionResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return DevActionResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEventsResponse:
		msg := &deadlock.CMsgClientToGCRecordClientEventsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return RecordClientEventsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgressResponse:
		msg := &deadlock.CMsgClientToGCSetNewPlayerProgressResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SetNewPlayerProgressResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSyncResponse:
		msg := &deadlock.CMsgClientToGCUpdateAccountSyncResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return UpdateAccountSyncResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoiceResponse:
		msg := &deadlock.CMsgClientToGCGetHeroChoiceResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetHeroChoiceResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHeroResponse:
		msg := &deadlock.CMsgClientToGCUnlockHeroResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return UnlockHeroResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlockResponse:
		msg := &deadlock.CMsgClientToGCBookUnlockResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return BookUnlockResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetBookResponse:
		msg := &deadlock.CMsgClientToGCGetBookResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetBookResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUserResponse:
		msg := &deadlock.CMsgClientToGCSubmitPlaytestUserResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SubmitPlaytestUserResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildResponse:
		msg := &deadlock.CMsgClientToGCUpdateHeroBuildResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return UpdateHeroBuildResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuildsResponse:
		msg := &deadlock.CMsgClientToGCFindHeroBuildsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return FindHeroBuildsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatchResponse:
		msg := &deadlock.CMsgClientToGCReportPlayerFromMatchResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return ReportPlayerFromMatchResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReportsResponse:
		msg := &deadlock.CMsgClientToGCGetAccountMatchReportsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetAccountMatchReportsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuildResponse:
		msg := &deadlock.CMsgClientToGCDeleteHeroBuildResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return DeleteHeroBuildResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatchesResponse:
		msg := &deadlock.CMsgClientToGCGetActiveMatchesResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetActiveMatchesResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLinkResponse:
		msg := &deadlock.CMsgClientToGCGetDiscordLinkResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetDiscordLinkResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCPartySetModeResponse:
		msg := &deadlock.CMsgClientToGCPartySetModeResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return PartySetModeResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccessResponse:
		msg := &deadlock.CMsgClientToGCGrantForumAccessResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GrantForumAccessResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequestResponse:
		msg := &deadlock.CMsgClientToGCModeratorRequestResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return ModeratorRequestResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatusResponse:
		msg := &deadlock.CMsgClientToGCGetFriendGameStatusResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetFriendGameStatusResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreferenceResponse:
		msg := &deadlock.CMsgClientToGCUpdateHeroBuildPreferenceResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return UpdateHeroBuildPreferenceResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetOldHeroBuildDataResponse:
		msg := &deadlock.CMsgClientToGCGetOldHeroBuildDataResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetOldHeroBuildDataResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatchResponse:
		msg := &deadlock.CMsgClientToGCCommendPlayerFromMatchResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return CommendPlayerFromMatchResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankingsResponse:
		msg := &deadlock.CMsgClientToGCGetHeroMMRRankingsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetHeroMMRRankingsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboardResponse:
		msg := &deadlock.CMsgClientToGCGetLeaderboardResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetLeaderboardResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboardsResponse:
		msg := &deadlock.CMsgClientToGCGetAccountLeaderboardsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return GetAccountLeaderboardsResponse{msg}, nil

	case deadlock.EGCCitadelClientMessages_k_EMsgClientToGCSetHideHolidayModelsResponse:
		msg := &deadlock.CMsgClientToGCSetHideHolidayModelsResponse{}
		err := proto.Unmarshal(packet.Body, msg)
		if err != nil {
			return nil, err
		}
		return SetHideHolidayModelsResponse{msg}, nil
	}
	return nil, fmt.Errorf("unknown message type: %d", packet.MsgType)
}
