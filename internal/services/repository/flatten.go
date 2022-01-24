package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/datadrivers/terraform-provider-nexus/internal/tools"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenCleanup(cleanup *repository.Cleanup) []map[string]interface{} {
	if cleanup == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"policy_names": tools.StringSliceToInterfaceSlice(cleanup.PolicyNames),
		},
	}
}

func flattenComponent(component *repository.Component) []map[string]interface{} {
	if component == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"proprietary_components": component.ProprietaryComponents,
		},
	}
}

func flattenGroup(group *repository.Group) []map[string]interface{} {
	if group == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"member_names": tools.StringSliceToInterfaceSlice(group.MemberNames),
		},
	}
}

func flattenHTTPClient(httpClient *repository.HTTPClient, d *schema.ResourceData) []map[string]interface{} {
	if httpClient == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"authentication": flattenHTTPClientAuthentication(httpClient.Authentication, d),
			"auto_block":     httpClient.AutoBlock,
			"blocked":        httpClient.Blocked,
			"connection":     flattenHTTPClientConnection(httpClient.Connection),
		},
	}
}

func flattenHTTPClientAuthentication(auth *repository.HTTPClientAuthentication, d *schema.ResourceData) []map[string]interface{} {
	if auth == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"ntlm_domain": auth.NTLMDomain,
			"ntlm_host":   auth.NTLMHost,
			"type":        auth.Type,
			"username":    auth.Username,
			"password":    d.Get("http_client.0.authentication.0.password").(string),
		},
	}
}

func flattenHTTPClientConnection(conn *repository.HTTPClientConnection) []map[string]interface{} {
	if conn == nil {
		return nil
	}
	data := map[string]interface{}{
		"user_agent_suffix": conn.UserAgentSuffix,
	}
	if conn.EnableCircularRedirects != nil {
		data["enable_circular_redirects"] = *conn.EnableCircularRedirects
	}
	if conn.EnableCookies != nil {
		data["enable_cookies"] = *conn.EnableCookies
	}
	if conn.Retries != nil {
		data["retries"] = *conn.Retries
	}
	if conn.Timeout != nil {
		data["timeout"] = *conn.Timeout
	}
	if conn.UseTrustStore != nil {
		data["use_trust_store"] = *conn.UseTrustStore
	}
	return []map[string]interface{}{data}
}

func flattenNegativeCache(negativeCache *repository.NegativeCache) []map[string]interface{} {
	if negativeCache == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"enabled": negativeCache.Enabled,
			"ttl":     negativeCache.TTL,
		},
	}
}

func flattenProxy(proxy *repository.Proxy) []map[string]interface{} {
	if proxy == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"content_max_age":  proxy.ContentMaxAge,
			"metadata_max_age": proxy.MetadataMaxAge,
			"remote_url":       proxy.RemoteURL,
		},
	}
}

func flattenStorage(storage *repository.Storage, d *schema.ResourceData) []map[string]interface{} {
	if storage == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"blob_store_name":                storage.BlobStoreName,
			"strict_content_type_validation": storage.StrictContentTypeValidation,
		},
	}
}

func flattenHostedStorage(storage *repository.HostedStorage, d *schema.ResourceData) []map[string]interface{} {
	if storage == nil {
		return nil
	}
	data := map[string]interface{}{
		"blob_store_name":                storage.BlobStoreName,
		"strict_content_type_validation": storage.StrictContentTypeValidation,
	}
	if storage.WritePolicy != nil {
		data["write_policy"] = storage.WritePolicy
	}
	return []map[string]interface{}{data}
}
