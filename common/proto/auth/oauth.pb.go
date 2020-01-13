// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OAuth2ClientConfig struct {
	ClientID      string   `protobuf:"bytes,1,opt,name=ClientID,json=client_id" json:"ClientID,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=Name,json=client_name" json:"Name,omitempty"`
	Secret        string   `protobuf:"bytes,3,opt,name=Secret,json=client_secret" json:"Secret,omitempty"`
	RedirectURIs  []string `protobuf:"bytes,4,rep,name=RedirectURIs,json=redirect_uris" json:"RedirectURIs,omitempty"`
	GrantTypes    []string `protobuf:"bytes,5,rep,name=GrantTypes,json=grant_types" json:"GrantTypes,omitempty"`
	ResponseTypes []string `protobuf:"bytes,6,rep,name=ResponseTypes,json=response_types" json:"ResponseTypes,omitempty"`
	Scope         string   `protobuf:"bytes,7,opt,name=Scope,json=scope" json:"Scope,omitempty"`
	Audience      []string `protobuf:"bytes,8,rep,name=Audience,json=audience" json:"Audience,omitempty"`
}

func (m *OAuth2ClientConfig) Reset()                    { *m = OAuth2ClientConfig{} }
func (m *OAuth2ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ClientConfig) ProtoMessage()               {}
func (*OAuth2ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *OAuth2ClientConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ClientConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuth2ClientConfig) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OAuth2ClientConfig) GetRedirectURIs() []string {
	if m != nil {
		return m.RedirectURIs
	}
	return nil
}

func (m *OAuth2ClientConfig) GetGrantTypes() []string {
	if m != nil {
		return m.GrantTypes
	}
	return nil
}

func (m *OAuth2ClientConfig) GetResponseTypes() []string {
	if m != nil {
		return m.ResponseTypes
	}
	return nil
}

func (m *OAuth2ClientConfig) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *OAuth2ClientConfig) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

type OAuth2ConnectorCollection struct {
	Connectors []*OAuth2ConnectorConfig `protobuf:"bytes,1,rep,name=connectors" json:"connectors,omitempty"`
}

func (m *OAuth2ConnectorCollection) Reset()                    { *m = OAuth2ConnectorCollection{} }
func (m *OAuth2ConnectorCollection) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorCollection) ProtoMessage()               {}
func (*OAuth2ConnectorCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *OAuth2ConnectorCollection) GetConnectors() []*OAuth2ConnectorConfig {
	if m != nil {
		return m.Connectors
	}
	return nil
}

type OAuth2ConnectorConfig struct {
	Id           string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type         string                 `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	MappingRules []*OAuth2MappingRule   `protobuf:"bytes,4,rep,name=mappingRules" json:"mappingRules,omitempty"`
	Config       *google_protobuf.Value `protobuf:"bytes,5,opt,name=config" json:"config,omitempty"`
}

func (m *OAuth2ConnectorConfig) Reset()                    { *m = OAuth2ConnectorConfig{} }
func (m *OAuth2ConnectorConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorConfig) ProtoMessage()               {}
func (*OAuth2ConnectorConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *OAuth2ConnectorConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OAuth2ConnectorConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OAuth2ConnectorConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuth2ConnectorConfig) GetMappingRules() []*OAuth2MappingRule {
	if m != nil {
		return m.MappingRules
	}
	return nil
}

func (m *OAuth2ConnectorConfig) GetConfig() *google_protobuf.Value {
	if m != nil {
		return m.Config
	}
	return nil
}

type OAuth2MappingRule struct {
	LeftAttribute  string `protobuf:"bytes,1,opt,name=LeftAttribute" json:"LeftAttribute,omitempty"`
	RuleString     string `protobuf:"bytes,2,opt,name=RuleString" json:"RuleString,omitempty"`
	RightAttribute string `protobuf:"bytes,3,opt,name=RightAttribute" json:"RightAttribute,omitempty"`
}

func (m *OAuth2MappingRule) Reset()                    { *m = OAuth2MappingRule{} }
func (m *OAuth2MappingRule) String() string            { return proto.CompactTextString(m) }
func (*OAuth2MappingRule) ProtoMessage()               {}
func (*OAuth2MappingRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *OAuth2MappingRule) GetLeftAttribute() string {
	if m != nil {
		return m.LeftAttribute
	}
	return ""
}

func (m *OAuth2MappingRule) GetRuleString() string {
	if m != nil {
		return m.RuleString
	}
	return ""
}

func (m *OAuth2MappingRule) GetRightAttribute() string {
	if m != nil {
		return m.RightAttribute
	}
	return ""
}

type OAuth2ConnectorPydioConfig struct {
	Pydioconnectors []*OAuth2ConnectorPydioConfig_Connector `protobuf:"bytes,1,rep,name=pydioconnectors" json:"pydioconnectors,omitempty"`
}

func (m *OAuth2ConnectorPydioConfig) Reset()                    { *m = OAuth2ConnectorPydioConfig{} }
func (m *OAuth2ConnectorPydioConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorPydioConfig) ProtoMessage()               {}
func (*OAuth2ConnectorPydioConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *OAuth2ConnectorPydioConfig) GetPydioconnectors() []*OAuth2ConnectorPydioConfig_Connector {
	if m != nil {
		return m.Pydioconnectors
	}
	return nil
}

type OAuth2ConnectorPydioConfig_Connector struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *OAuth2ConnectorPydioConfig_Connector) Reset()         { *m = OAuth2ConnectorPydioConfig_Connector{} }
func (m *OAuth2ConnectorPydioConfig_Connector) String() string { return proto.CompactTextString(m) }
func (*OAuth2ConnectorPydioConfig_Connector) ProtoMessage()    {}
func (*OAuth2ConnectorPydioConfig_Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{4, 0}
}

func (m *OAuth2ConnectorPydioConfig_Connector) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OAuth2ConnectorPydioConfig_Connector) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuth2ConnectorPydioConfig_Connector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type OAuth2ConnectorOIDCConfig struct {
	Issuer                    string   `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	ClientID                  string   `protobuf:"bytes,2,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret              string   `protobuf:"bytes,3,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI               string   `protobuf:"bytes,4,opt,name=redirectURI" json:"redirectURI,omitempty"`
	BasicAuthUnsupported      bool     `protobuf:"varint,5,opt,name=basicAuthUnsupported" json:"basicAuthUnsupported,omitempty"`
	HostedDomains             []string `protobuf:"bytes,6,rep,name=hostedDomains" json:"hostedDomains,omitempty"`
	Scopes                    []string `protobuf:"bytes,7,rep,name=scopes" json:"scopes,omitempty"`
	InsecureSkipEmailVerified bool     `protobuf:"varint,8,opt,name=insecureSkipEmailVerified" json:"insecureSkipEmailVerified,omitempty"`
	GetUserInfo               bool     `protobuf:"varint,9,opt,name=getUserInfo" json:"getUserInfo,omitempty"`
	UserIDKey                 string   `protobuf:"bytes,10,opt,name=userIDKey" json:"userIDKey,omitempty"`
	UserNameKey               string   `protobuf:"bytes,11,opt,name=userNameKey" json:"userNameKey,omitempty"`
}

func (m *OAuth2ConnectorOIDCConfig) Reset()                    { *m = OAuth2ConnectorOIDCConfig{} }
func (m *OAuth2ConnectorOIDCConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorOIDCConfig) ProtoMessage()               {}
func (*OAuth2ConnectorOIDCConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *OAuth2ConnectorOIDCConfig) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *OAuth2ConnectorOIDCConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorOIDCConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorOIDCConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorOIDCConfig) GetBasicAuthUnsupported() bool {
	if m != nil {
		return m.BasicAuthUnsupported
	}
	return false
}

func (m *OAuth2ConnectorOIDCConfig) GetHostedDomains() []string {
	if m != nil {
		return m.HostedDomains
	}
	return nil
}

func (m *OAuth2ConnectorOIDCConfig) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *OAuth2ConnectorOIDCConfig) GetInsecureSkipEmailVerified() bool {
	if m != nil {
		return m.InsecureSkipEmailVerified
	}
	return false
}

func (m *OAuth2ConnectorOIDCConfig) GetGetUserInfo() bool {
	if m != nil {
		return m.GetUserInfo
	}
	return false
}

func (m *OAuth2ConnectorOIDCConfig) GetUserIDKey() string {
	if m != nil {
		return m.UserIDKey
	}
	return ""
}

func (m *OAuth2ConnectorOIDCConfig) GetUserNameKey() string {
	if m != nil {
		return m.UserNameKey
	}
	return ""
}

type OAuth2ConnectorSAMLConfig struct {
	SsoURL                          string `protobuf:"bytes,1,opt,name=ssoURL" json:"ssoURL,omitempty"`
	Ca                              string `protobuf:"bytes,2,opt,name=ca" json:"ca,omitempty"`
	RedirectURI                     string `protobuf:"bytes,3,opt,name=redirectURI" json:"redirectURI,omitempty"`
	UsernameAttr                    string `protobuf:"bytes,4,opt,name=usernameAttr" json:"usernameAttr,omitempty"`
	EmailAttr                       string `protobuf:"bytes,5,opt,name=emailAttr" json:"emailAttr,omitempty"`
	GroupsAttr                      string `protobuf:"bytes,6,opt,name=groupsAttr" json:"groupsAttr,omitempty"`
	CaData                          string `protobuf:"bytes,7,opt,name=caData" json:"caData,omitempty"`
	InsecureSkipSignatureValidation bool   `protobuf:"varint,8,opt,name=insecureSkipSignatureValidation" json:"insecureSkipSignatureValidation,omitempty"`
	EntityIssuer                    string `protobuf:"bytes,9,opt,name=entityIssuer" json:"entityIssuer,omitempty"`
	SsoIssuer                       string `protobuf:"bytes,10,opt,name=ssoIssuer" json:"ssoIssuer,omitempty"`
	GroupsDelim                     string `protobuf:"bytes,11,opt,name=groupsDelim" json:"groupsDelim,omitempty"`
	NameIDPolicyFormat              string `protobuf:"bytes,12,opt,name=nameIDPolicyFormat" json:"nameIDPolicyFormat,omitempty"`
}

func (m *OAuth2ConnectorSAMLConfig) Reset()                    { *m = OAuth2ConnectorSAMLConfig{} }
func (m *OAuth2ConnectorSAMLConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorSAMLConfig) ProtoMessage()               {}
func (*OAuth2ConnectorSAMLConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *OAuth2ConnectorSAMLConfig) GetSsoURL() string {
	if m != nil {
		return m.SsoURL
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetCa() string {
	if m != nil {
		return m.Ca
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetUsernameAttr() string {
	if m != nil {
		return m.UsernameAttr
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetEmailAttr() string {
	if m != nil {
		return m.EmailAttr
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetGroupsAttr() string {
	if m != nil {
		return m.GroupsAttr
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetCaData() string {
	if m != nil {
		return m.CaData
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetInsecureSkipSignatureValidation() bool {
	if m != nil {
		return m.InsecureSkipSignatureValidation
	}
	return false
}

func (m *OAuth2ConnectorSAMLConfig) GetEntityIssuer() string {
	if m != nil {
		return m.EntityIssuer
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetSsoIssuer() string {
	if m != nil {
		return m.SsoIssuer
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetGroupsDelim() string {
	if m != nil {
		return m.GroupsDelim
	}
	return ""
}

func (m *OAuth2ConnectorSAMLConfig) GetNameIDPolicyFormat() string {
	if m != nil {
		return m.NameIDPolicyFormat
	}
	return ""
}

type OAuth2ConnectorBitbucketConfig struct {
	ClientID     string   `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret string   `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI  string   `protobuf:"bytes,3,opt,name=redirectURI" json:"redirectURI,omitempty"`
	Teams        []string `protobuf:"bytes,4,rep,name=teams" json:"teams,omitempty"`
}

func (m *OAuth2ConnectorBitbucketConfig) Reset()                    { *m = OAuth2ConnectorBitbucketConfig{} }
func (m *OAuth2ConnectorBitbucketConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorBitbucketConfig) ProtoMessage()               {}
func (*OAuth2ConnectorBitbucketConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *OAuth2ConnectorBitbucketConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorBitbucketConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorBitbucketConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorBitbucketConfig) GetTeams() []string {
	if m != nil {
		return m.Teams
	}
	return nil
}

type OAuth2ConnectorGithubConfig struct {
	ClientID      string                            `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret  string                            `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI   string                            `protobuf:"bytes,3,opt,name=redirectURI" json:"redirectURI,omitempty"`
	Orgs          []*OAuth2ConnectorGithubConfigOrg `protobuf:"bytes,4,rep,name=orgs" json:"orgs,omitempty"`
	LoadAllGroups bool                              `protobuf:"varint,5,opt,name=loadAllGroups" json:"loadAllGroups,omitempty"`
	TeamNameField string                            `protobuf:"bytes,6,opt,name=teamNameField" json:"teamNameField,omitempty"`
	UseLoginAsID  bool                              `protobuf:"varint,7,opt,name=useLoginAsID" json:"useLoginAsID,omitempty"`
	// For GitHub enterprise
	HostName string `protobuf:"bytes,8,opt,name=hostName" json:"hostName,omitempty"`
	RootCA   string `protobuf:"bytes,9,opt,name=rootCA" json:"rootCA,omitempty"`
}

func (m *OAuth2ConnectorGithubConfig) Reset()                    { *m = OAuth2ConnectorGithubConfig{} }
func (m *OAuth2ConnectorGithubConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorGithubConfig) ProtoMessage()               {}
func (*OAuth2ConnectorGithubConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *OAuth2ConnectorGithubConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfig) GetOrgs() []*OAuth2ConnectorGithubConfigOrg {
	if m != nil {
		return m.Orgs
	}
	return nil
}

func (m *OAuth2ConnectorGithubConfig) GetLoadAllGroups() bool {
	if m != nil {
		return m.LoadAllGroups
	}
	return false
}

func (m *OAuth2ConnectorGithubConfig) GetTeamNameField() string {
	if m != nil {
		return m.TeamNameField
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfig) GetUseLoginAsID() bool {
	if m != nil {
		return m.UseLoginAsID
	}
	return false
}

func (m *OAuth2ConnectorGithubConfig) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfig) GetRootCA() string {
	if m != nil {
		return m.RootCA
	}
	return ""
}

type OAuth2ConnectorGithubConfigOrg struct {
	Name  string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Teams []string `protobuf:"bytes,2,rep,name=teams" json:"teams,omitempty"`
}

func (m *OAuth2ConnectorGithubConfigOrg) Reset()                    { *m = OAuth2ConnectorGithubConfigOrg{} }
func (m *OAuth2ConnectorGithubConfigOrg) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorGithubConfigOrg) ProtoMessage()               {}
func (*OAuth2ConnectorGithubConfigOrg) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *OAuth2ConnectorGithubConfigOrg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuth2ConnectorGithubConfigOrg) GetTeams() []string {
	if m != nil {
		return m.Teams
	}
	return nil
}

type OAuth2ConnectorGitlabConfig struct {
	BaseURL       string   `protobuf:"bytes,1,opt,name=baseURL" json:"baseURL,omitempty"`
	ClientID      string   `protobuf:"bytes,2,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret  string   `protobuf:"bytes,3,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI   string   `protobuf:"bytes,4,opt,name=redirectURI" json:"redirectURI,omitempty"`
	Groups        []string `protobuf:"bytes,5,rep,name=groups" json:"groups,omitempty"`
	UserLoginAsID bool     `protobuf:"varint,6,opt,name=userLoginAsID" json:"userLoginAsID,omitempty"`
}

func (m *OAuth2ConnectorGitlabConfig) Reset()                    { *m = OAuth2ConnectorGitlabConfig{} }
func (m *OAuth2ConnectorGitlabConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorGitlabConfig) ProtoMessage()               {}
func (*OAuth2ConnectorGitlabConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *OAuth2ConnectorGitlabConfig) GetBaseURL() string {
	if m != nil {
		return m.BaseURL
	}
	return ""
}

func (m *OAuth2ConnectorGitlabConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorGitlabConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorGitlabConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorGitlabConfig) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *OAuth2ConnectorGitlabConfig) GetUserLoginAsID() bool {
	if m != nil {
		return m.UserLoginAsID
	}
	return false
}

type OAuth2ConnectorLinkedinConfig struct {
	ClientID     string `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI  string `protobuf:"bytes,3,opt,name=redirectURI" json:"redirectURI,omitempty"`
}

func (m *OAuth2ConnectorLinkedinConfig) Reset()                    { *m = OAuth2ConnectorLinkedinConfig{} }
func (m *OAuth2ConnectorLinkedinConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorLinkedinConfig) ProtoMessage()               {}
func (*OAuth2ConnectorLinkedinConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *OAuth2ConnectorLinkedinConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorLinkedinConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorLinkedinConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

type OAuth2ConnectorMicrosoftConfig struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	RedirectURI          string   `protobuf:"bytes,3,opt,name=redirectURI" json:"redirectURI,omitempty"`
	Tenant               string   `protobuf:"bytes,4,opt,name=tenant" json:"tenant,omitempty"`
	Groups               []string `protobuf:"bytes,5,rep,name=groups" json:"groups,omitempty"`
	OnlySecurityGroups   bool     `protobuf:"varint,6,opt,name=onlySecurityGroups" json:"onlySecurityGroups,omitempty"`
	GroupNameFormat      string   `protobuf:"bytes,7,opt,name=groupNameFormat" json:"groupNameFormat,omitempty"`
	UseGroupsAsWhitelist bool     `protobuf:"varint,8,opt,name=useGroupsAsWhitelist" json:"useGroupsAsWhitelist,omitempty"`
}

func (m *OAuth2ConnectorMicrosoftConfig) Reset()                    { *m = OAuth2ConnectorMicrosoftConfig{} }
func (m *OAuth2ConnectorMicrosoftConfig) String() string            { return proto.CompactTextString(m) }
func (*OAuth2ConnectorMicrosoftConfig) ProtoMessage()               {}
func (*OAuth2ConnectorMicrosoftConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *OAuth2ConnectorMicrosoftConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OAuth2ConnectorMicrosoftConfig) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *OAuth2ConnectorMicrosoftConfig) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *OAuth2ConnectorMicrosoftConfig) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *OAuth2ConnectorMicrosoftConfig) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *OAuth2ConnectorMicrosoftConfig) GetOnlySecurityGroups() bool {
	if m != nil {
		return m.OnlySecurityGroups
	}
	return false
}

func (m *OAuth2ConnectorMicrosoftConfig) GetGroupNameFormat() string {
	if m != nil {
		return m.GroupNameFormat
	}
	return ""
}

func (m *OAuth2ConnectorMicrosoftConfig) GetUseGroupsAsWhitelist() bool {
	if m != nil {
		return m.UseGroupsAsWhitelist
	}
	return false
}

func init() {
	proto.RegisterType((*OAuth2ClientConfig)(nil), "auth.OAuth2ClientConfig")
	proto.RegisterType((*OAuth2ConnectorCollection)(nil), "auth.OAuth2ConnectorCollection")
	proto.RegisterType((*OAuth2ConnectorConfig)(nil), "auth.OAuth2ConnectorConfig")
	proto.RegisterType((*OAuth2MappingRule)(nil), "auth.OAuth2MappingRule")
	proto.RegisterType((*OAuth2ConnectorPydioConfig)(nil), "auth.OAuth2ConnectorPydioConfig")
	proto.RegisterType((*OAuth2ConnectorPydioConfig_Connector)(nil), "auth.OAuth2ConnectorPydioConfig.Connector")
	proto.RegisterType((*OAuth2ConnectorOIDCConfig)(nil), "auth.OAuth2ConnectorOIDCConfig")
	proto.RegisterType((*OAuth2ConnectorSAMLConfig)(nil), "auth.OAuth2ConnectorSAMLConfig")
	proto.RegisterType((*OAuth2ConnectorBitbucketConfig)(nil), "auth.OAuth2ConnectorBitbucketConfig")
	proto.RegisterType((*OAuth2ConnectorGithubConfig)(nil), "auth.OAuth2ConnectorGithubConfig")
	proto.RegisterType((*OAuth2ConnectorGithubConfigOrg)(nil), "auth.OAuth2ConnectorGithubConfigOrg")
	proto.RegisterType((*OAuth2ConnectorGitlabConfig)(nil), "auth.OAuth2ConnectorGitlabConfig")
	proto.RegisterType((*OAuth2ConnectorLinkedinConfig)(nil), "auth.OAuth2ConnectorLinkedinConfig")
	proto.RegisterType((*OAuth2ConnectorMicrosoftConfig)(nil), "auth.OAuth2ConnectorMicrosoftConfig")
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4b, 0x6e, 0x23, 0x37,
	0x13, 0x86, 0x9e, 0x23, 0x95, 0x3c, 0x1e, 0xfc, 0x84, 0x7f, 0xa7, 0x6d, 0xcf, 0xc3, 0xe8, 0x4c,
	0x02, 0x23, 0x0b, 0x0d, 0xe0, 0x6c, 0x02, 0x4c, 0x36, 0x8a, 0x94, 0x71, 0x94, 0xd8, 0xf1, 0xa0,
	0x35, 0x76, 0xb2, 0x33, 0xa8, 0x6e, 0xaa, 0x45, 0xb8, 0x45, 0x0a, 0x24, 0x7b, 0xa1, 0x65, 0xb6,
	0x39, 0x41, 0x6e, 0x90, 0x23, 0x64, 0x93, 0x00, 0xb9, 0x45, 0x56, 0xb9, 0x4b, 0x50, 0x24, 0x5b,
	0x6a, 0x3d, 0x06, 0xde, 0x79, 0xc7, 0xfa, 0xaa, 0xa8, 0xae, 0xfa, 0xbe, 0xaa, 0xa2, 0xa0, 0x23,
	0x69, 0x6e, 0xa6, 0xdd, 0xb9, 0x92, 0x46, 0x92, 0x3a, 0x9e, 0x8f, 0x9f, 0xa7, 0x52, 0xa6, 0x19,
	0x7b, 0x63, 0xb1, 0x71, 0x3e, 0x79, 0xa3, 0x8d, 0xca, 0x63, 0xe3, 0x62, 0xc2, 0x5f, 0xab, 0x40,
	0xae, 0x7b, 0xb9, 0x99, 0x9e, 0xf7, 0x33, 0xce, 0x84, 0xe9, 0x4b, 0x31, 0xe1, 0x29, 0x39, 0x81,
	0x96, 0xb3, 0x87, 0x83, 0xa0, 0x72, 0x5a, 0x39, 0x6b, 0x47, 0xed, 0xd8, 0xda, 0x77, 0x3c, 0x21,
	0x47, 0x50, 0xff, 0x91, 0xce, 0x58, 0x50, 0xb5, 0x8e, 0x8e, 0x77, 0x08, 0x3a, 0x63, 0xe4, 0x05,
	0x34, 0x47, 0x2c, 0x56, 0xcc, 0x04, 0x35, 0xeb, 0x7c, 0xea, 0x9d, 0xda, 0x82, 0xe4, 0x53, 0xd8,
	0x8b, 0x58, 0xc2, 0x15, 0x8b, 0xcd, 0x4d, 0x34, 0xd4, 0x41, 0xfd, 0xb4, 0x86, 0x41, 0xca, 0x63,
	0x77, 0xb9, 0xe2, 0x9a, 0xbc, 0x02, 0xb8, 0x50, 0x54, 0x98, 0x0f, 0x8b, 0x39, 0xd3, 0x41, 0xc3,
	0x86, 0x74, 0x52, 0x44, 0xee, 0x0c, 0x42, 0xe4, 0x33, 0x78, 0x1a, 0x31, 0x3d, 0x97, 0x42, 0x33,
	0x17, 0xd3, 0xb4, 0x31, 0xfb, 0xca, 0x83, 0x3e, 0xec, 0x00, 0x1a, 0xa3, 0x58, 0xce, 0x59, 0xf0,
	0xc4, 0xa6, 0xd2, 0xd0, 0x68, 0x90, 0x63, 0x68, 0xf5, 0xf2, 0x84, 0x33, 0x11, 0xb3, 0xa0, 0x65,
	0xef, 0xb5, 0xa8, 0xb7, 0xc3, 0x9f, 0xe1, 0xc8, 0x73, 0x21, 0x85, 0x60, 0xb1, 0x91, 0xaa, 0x2f,
	0xb3, 0x8c, 0xc5, 0x86, 0x4b, 0x41, 0xde, 0x02, 0xc4, 0x05, 0xac, 0x83, 0xca, 0x69, 0xed, 0xac,
	0x73, 0x7e, 0xd2, 0xb5, 0x74, 0x6f, 0x5d, 0x42, 0x0e, 0xa3, 0x52, 0x78, 0xf8, 0x67, 0x05, 0xfe,
	0xbf, 0x33, 0x8a, 0xec, 0x43, 0x95, 0x27, 0x9e, 0xe3, 0x2a, 0x4f, 0x08, 0x81, 0x3a, 0xa6, 0xef,
	0xc9, 0xb5, 0x67, 0xc4, 0x90, 0x5d, 0xcf, 0xa9, 0x3d, 0x93, 0xb7, 0xb0, 0x37, 0xa3, 0xf3, 0x39,
	0x17, 0x69, 0x94, 0x67, 0xcc, 0x51, 0xd9, 0x39, 0xff, 0xa4, 0x9c, 0xd0, 0xd5, 0xca, 0x1f, 0xad,
	0x05, 0x93, 0x2e, 0x34, 0x63, 0xfb, 0xf9, 0xa0, 0x71, 0x5a, 0x39, 0xeb, 0x9c, 0x1f, 0x76, 0x5d,
	0x93, 0x74, 0x8b, 0x26, 0xe9, 0xde, 0xd2, 0x2c, 0x67, 0x91, 0x8f, 0x0a, 0x7f, 0xa9, 0xc0, 0xff,
	0xb6, 0x7e, 0x93, 0xbc, 0x86, 0xa7, 0x97, 0x6c, 0x62, 0x7a, 0xc6, 0x28, 0x3e, 0xce, 0x0d, 0xf3,
	0x55, 0xac, 0x83, 0xe4, 0x25, 0x00, 0x46, 0x8f, 0x8c, 0xe2, 0x22, 0xf5, 0x65, 0x95, 0x10, 0xf2,
	0x39, 0xec, 0x47, 0x3c, 0x9d, 0x96, 0x7e, 0xc6, 0x95, 0xb9, 0x81, 0x86, 0x7f, 0x54, 0xe0, 0x78,
	0x83, 0xc2, 0xf7, 0x8b, 0x84, 0x4b, 0xcf, 0xe3, 0x07, 0x78, 0x36, 0x47, 0x73, 0x4b, 0xa3, 0x2f,
	0x76, 0x6a, 0x54, 0xba, 0xda, 0x5d, 0x82, 0xd1, 0xe6, 0x4f, 0x1c, 0xf7, 0xa1, 0xbd, 0xf4, 0x96,
	0xa4, 0x6a, 0x14, 0x52, 0x89, 0xd5, 0x1c, 0x38, 0x59, 0x0a, 0xf9, 0x6a, 0x2b, 0xf9, 0xc2, 0xdf,
	0x6b, 0x5b, 0x7d, 0x75, 0x3d, 0x1c, 0xf4, 0x7d, 0xe2, 0x87, 0xd0, 0xe4, 0x5a, 0xe7, 0x4c, 0x79,
	0xfa, 0xbc, 0x85, 0x8d, 0x1a, 0x17, 0x23, 0xe8, 0xbe, 0xb0, 0xb4, 0x49, 0x08, 0x7b, 0xee, 0xbc,
	0x36, 0x6c, 0x6b, 0x18, 0x39, 0x85, 0x8e, 0x5a, 0xcd, 0x5a, 0x50, 0x77, 0xc3, 0x5a, 0x82, 0xc8,
	0x39, 0x1c, 0x8c, 0xa9, 0xe6, 0x31, 0xa6, 0x76, 0x23, 0x74, 0x3e, 0x9f, 0x4b, 0x65, 0x58, 0x62,
	0x7b, 0xa2, 0x15, 0xed, 0xf4, 0xa1, 0xe6, 0x53, 0xa9, 0x0d, 0x4b, 0x06, 0x72, 0x46, 0xb9, 0x28,
	0x66, 0x6f, 0x1d, 0xc4, 0x9a, 0xec, 0xb4, 0xe9, 0xe0, 0x89, 0x75, 0x7b, 0x8b, 0x7c, 0x0d, 0x47,
	0x5c, 0x68, 0x16, 0xe7, 0x8a, 0x8d, 0xee, 0xf9, 0xfc, 0xdb, 0x19, 0xe5, 0xd9, 0x2d, 0x53, 0x7c,
	0xc2, 0x59, 0x12, 0xb4, 0xec, 0x67, 0x3f, 0x1e, 0x80, 0x15, 0xa5, 0xcc, 0xdc, 0x68, 0xa6, 0x86,
	0x62, 0x22, 0x83, 0xb6, 0x8d, 0x2f, 0x43, 0xe4, 0x39, 0xb4, 0x73, 0x3c, 0x0f, 0x7e, 0x60, 0x8b,
	0x00, 0xdc, 0xde, 0x5a, 0x02, 0x78, 0x1f, 0x0d, 0xdc, 0x5d, 0xe8, 0xef, 0x38, 0x46, 0x4a, 0x50,
	0xf8, 0xd7, 0xb6, 0x52, 0xa3, 0xde, 0xd5, 0xe5, 0x4a, 0x29, 0xad, 0xe5, 0x4d, 0x74, 0x59, 0x28,
	0xe5, 0x2c, 0xec, 0x8b, 0x98, 0x7a, 0x8d, 0xaa, 0x31, 0xdd, 0x64, 0xbe, 0xb6, 0xcd, 0x7c, 0x08,
	0x7b, 0xf8, 0x59, 0xec, 0x18, 0x6c, 0x70, 0x2f, 0xce, 0x1a, 0x86, 0xb5, 0x30, 0x2c, 0xdf, 0x06,
	0x34, 0x5c, 0x2d, 0x4b, 0x00, 0xa7, 0x2a, 0x55, 0x32, 0x9f, 0x6b, 0xeb, 0x6e, 0xba, 0xa9, 0x5a,
	0x21, 0x98, 0x6b, 0x4c, 0x07, 0xd4, 0x50, 0xbf, 0xfd, 0xbc, 0x45, 0xbe, 0x83, 0x57, 0x65, 0x82,
	0x47, 0x3c, 0x15, 0xd4, 0xe4, 0x8a, 0xdd, 0xd2, 0x8c, 0x27, 0x14, 0x17, 0x9d, 0xd7, 0xe1, 0xa1,
	0x30, 0xac, 0x81, 0x09, 0xc3, 0xcd, 0x62, 0xe8, 0xba, 0xb7, 0xed, 0x6a, 0x28, 0x63, 0x58, 0x83,
	0xd6, 0xd2, 0x07, 0x78, 0x3d, 0x96, 0x80, 0xd5, 0xd3, 0x66, 0x3c, 0x60, 0x19, 0x9f, 0x15, 0x7a,
	0x94, 0x20, 0xd2, 0x05, 0x82, 0x7c, 0x0c, 0x07, 0xef, 0x65, 0xc6, 0xe3, 0xc5, 0x3b, 0xa9, 0x66,
	0xd4, 0x04, 0x7b, 0x36, 0x70, 0x87, 0x27, 0xfc, 0xad, 0x02, 0x2f, 0x37, 0xf4, 0xfb, 0x86, 0x9b,
	0x71, 0x1e, 0xdf, 0xb3, 0xe2, 0x65, 0x2b, 0x8f, 0x55, 0xe5, 0x81, 0xb1, 0xaa, 0x3e, 0x3c, 0x56,
	0x3b, 0xc4, 0x3d, 0x80, 0x86, 0x61, 0x74, 0x56, 0xbc, 0x6e, 0xce, 0x08, 0xff, 0xad, 0xc2, 0xc9,
	0x46, 0x6a, 0x17, 0xdc, 0x4c, 0xf3, 0xf1, 0xa3, 0xe5, 0xf5, 0x15, 0xd4, 0xa5, 0x4a, 0x8b, 0x97,
	0xe2, 0xf5, 0xce, 0xb5, 0x58, 0x4e, 0xe9, 0x5a, 0xa5, 0x91, 0xbd, 0x81, 0x43, 0x9f, 0x49, 0x9a,
	0xf4, 0xb2, 0xec, 0xc2, 0x8a, 0xe3, 0x37, 0xc4, 0x3a, 0x88, 0x51, 0x58, 0x2a, 0xce, 0xd2, 0x3b,
	0xce, 0xb2, 0xc4, 0x77, 0xe5, 0x3a, 0xe8, 0x5b, 0xff, 0x52, 0xa6, 0x5c, 0xf4, 0xf4, 0x70, 0x60,
	0xdb, 0xb3, 0x15, 0xad, 0x61, 0xc8, 0x05, 0xee, 0x13, 0xfb, 0x27, 0xa3, 0xe5, 0xb8, 0x28, 0x6c,
	0x6c, 0x6c, 0x25, 0xa5, 0xe9, 0xf7, 0x7c, 0xc3, 0x79, 0x2b, 0xfc, 0x7e, 0x4b, 0xf9, 0x8d, 0x5a,
	0x96, 0xeb, 0xba, 0x52, 0x5a, 0xd7, 0x4b, 0xad, 0xaa, 0x65, 0xad, 0xfe, 0xa9, 0xec, 0xd2, 0x2a,
	0xa3, 0x85, 0x56, 0x01, 0x3c, 0x19, 0x53, 0xcd, 0x56, 0x9b, 0xa0, 0x30, 0x1f, 0x61, 0x69, 0x1f,
	0x42, 0x33, 0x2d, 0x44, 0xb0, 0xab, 0x35, 0x5d, 0xb2, 0x8f, 0xeb, 0x63, 0x45, 0x6c, 0xd3, 0x69,
	0xb4, 0x06, 0xe2, 0x43, 0xfe, 0x62, 0xa3, 0xb2, 0x4b, 0x2e, 0xee, 0x59, 0xc2, 0xc5, 0x63, 0xf5,
	0x61, 0xf8, 0x77, 0x75, 0x4b, 0xaa, 0x2b, 0x1e, 0x2b, 0xa9, 0xe5, 0xe4, 0xf1, 0x86, 0xf4, 0x10,
	0x9a, 0x86, 0x09, 0x2a, 0x8c, 0xe7, 0xd8, 0x5b, 0x1f, 0xa5, 0xb7, 0x0b, 0x44, 0x8a, 0x6c, 0x31,
	0xc2, 0x95, 0xc8, 0xcd, 0xc2, 0xcf, 0x81, 0xe3, 0x78, 0x87, 0x87, 0x9c, 0xc1, 0x33, 0x7b, 0xd3,
	0x36, 0xbe, 0x5b, 0x5b, 0x6e, 0x11, 0x6f, 0xc2, 0xf8, 0x0a, 0xe7, 0x9a, 0xb9, 0x6b, 0x3d, 0xfd,
	0xd3, 0x94, 0x1b, 0x96, 0x71, 0x6d, 0xfc, 0x1a, 0xde, 0xe9, 0x1b, 0x37, 0xed, 0xff, 0xb4, 0x2f,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x02, 0x9c, 0x69, 0xef, 0x0b, 0x00, 0x00,
}