package steamapi

import (
	"github.com/246859/steamapi/types/app"
	"github.com/246859/steamapi/types/steam"
	"net/http"
)

func (c *Client) ISteamApps() *ISteamApps {
	return &ISteamApps{c: c}
}

// ISteamApps see https://partner.steamgames.com/doc/webapi/ISteamApps
type ISteamApps struct {
	c *Client
}

// GetAppBeta see https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppBeta
func (i *ISteamApps) GetAppBeta(appId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(steam.AppId{AppId: appId}))
	return i.c.Unknown(http.MethodGet, PartnerHost, app.URLGetAppBetas, ops...)
}

// GetAppBuilds see https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppBuilds
func (i *ISteamApps) GetAppBuilds(buildsQueryOpt app.BuildsQueryOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(buildsQueryOpt))
	return i.c.Unknown(http.MethodGet, PartnerHost, app.URLGetAppBuilds, ops...)
}

// GetAppDepotVersion see https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppDepotVersion
func (i *ISteamApps) GetAppDepotVersion(appId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(steam.AppId{AppId: appId}))
	return i.c.Unknown(http.MethodGet, PartnerHost, app.URLGetAppDepotVersions, ops...)
}

// GetAppList see https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppList
func (i *ISteamApps) GetAppList(ops ...RequestOption) (app.PublicAppList, error) {
	var appList app.PublicAppList
	_, err := i.c.Get(PublicHost, app.URLGetAppList, &appList, ops...)
	if err != nil {
		return appList, err
	}
	return appList, nil
}

// GetPartnerAppListForWebAPIKey see https://partner.steamgames.com/doc/webapi/ISteamApps#GetPartnerAppListForWebAPIKey
func (i *ISteamApps) GetPartnerAppListForWebAPIKey(filter app.PartnerAppQueryFilter, ops ...RequestOption) (app.PartnerAppList, error) {
	var appList app.PartnerAppList
	ops = joinRequestOptions(ops, WithQueryForm(filter))
	_, err := i.c.Get(PartnerHost, app.URLGetPartnerAppListForWebAPIKey, &appList, ops...)
	if err != nil {
		return appList, err
	}
	return appList, nil
}

// GetPlayersBanned see https://partner.steamgames.com/doc/webapi/ISteamApps#GetPlayersBanned
func (i *ISteamApps) GetPlayersBanned(appId uint, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(steam.AppId{AppId: appId}))
	return i.c.Unknown(http.MethodGet, PartnerHost, app.URLGetPlayersBanned, ops...)
}

// GetServerList see https://partner.steamgames.com/doc/webapi/ISteamApps#GetServerList
func (i *ISteamApps) GetServerList(filter app.ServerListQueryFilter, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithQueryForm(filter))
	return i.c.Unknown(http.MethodGet, PartnerHost, app.URLServerList, ops...)
}

// GetServersAtAddress see https://partner.steamgames.com/doc/webapi/ISteamApps#GetServersAtAddress
func (i *ISteamApps) GetServersAtAddress(addr string, ops ...RequestOption) (app.ServerAddressList, error) {
	var servers app.ServerAddressList
	ops = joinRequestOptions(ops, WithQueryForm(app.ServerAddressQueryOption{Address: addr}))
	_, err := i.c.Get(PublicHost, app.URLGetServersAtAddress, &servers, ops...)
	if err != nil {
		return servers, err
	}
	return servers, nil
}

// SetAppBuildLive see https://partner.steamgames.com/doc/webapi/ISteamApps#SetAppBuildLive
func (i *ISteamApps) SetAppBuildLive(updateOption app.BuildLiveUpdateOption, ops ...RequestOption) (steam.CommonResponse, error) {
	ops = joinRequestOptions(ops, WithFormData(updateOption))
	return i.c.Unknown(http.MethodPost, PartnerHost, app.URLSetAppBuildLive, ops...)
}

// UpToDateCheck see https://partner.steamgames.com/doc/webapi/ISteamApps#UpToDateCheck
func (i *ISteamApps) UpToDateCheck(opt app.UpToDateCheckQueryOption, ops ...RequestOption) (app.UpToDateCheck, error) {
	var check app.UpToDateCheck
	_, err := i.c.Get(PublicHost, app.URLUpToDateCheck, &check, joinRequestOptions(ops, WithQueryForm(opt))...)
	if err != nil {
		return check, err
	}
	return check, nil
}
