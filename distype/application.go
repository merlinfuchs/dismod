package distype

type Application struct {
	ID                             Snowflake         `json:"id"`
	Name                           string            `json:"name"`
	Icon                           Nullable[string]  `json:"icon"`
	Description                    string            `json:"description"`
	RPCOrigins                     []string          `json:"rpc_origins,omitempty"`
	BotPublic                      bool              `json:"bot_public"`
	BotRequireCodeGrant            bool              `json:"bot_require_code_grant"`
	Bot                            *User             `json:"bot,omitempty"`
	TermsOfServiceURL              *string           `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL               *string           `json:"privacy_policy_url,omitempty"`
	Owner                          *User             `json:"owner,omitempty"`
	Summary                        string            `json:"summary"`
	VerifyKey                      string            `json:"verify_key"`
	Team                           Nullable[Team]    `json:"team"`
	GuildID                        *Snowflake        `json:"guild_id,omitempty"`
	Guild                          *Guild            `json:"guild,omitempty"`
	PrimarySKUID                   *Snowflake        `json:"primary_sku_id,omitempty"`
	Slug                           *string           `json:"slug,omitempty"`
	CoverImage                     *string           `json:"cover_image,omitempty"`
	Flags                          *ApplicationFlags `json:"flags,omitempty"`
	ApproximateGuildCount          *int              `json:"approximate_guild_count,omitempty"`
	RedirectURIs                   []string          `json:"redirect_uris,omitempty"`
	InteractionsEndpointURL        *string           `json:"interactions_endpoint_url,omitempty"`
	RoleConnectionsVerificationURL *string           `json:"role_connections_verification_url,omitempty"`
	Tags                           []string          `json:"tags,omitempty"`
	InstallParams                  *InstallParams    `json:"install_params,omitempty"`
	CustomInstallURL               *string           `json:"custom_install_url,omitempty"`
}

type ApplicationFlags int

const (
	ApplicationFlagsAutoModerationRuleCreateBadge ApplicationFlags = 1 << 6
	ApplicationFlagsGatewayPresence               ApplicationFlags = 1 << 12
	ApplicationFlagsGatewayPresenceLimited        ApplicationFlags = 1 << 13
	ApplicationFlagsGatewayGuildMembers           ApplicationFlags = 1 << 14
	ApplicationFlagsGatewayGuildMembersLimited    ApplicationFlags = 1 << 15
	ApplicationFlagsVerificationPendingGuildLimit ApplicationFlags = 1 << 16
	ApplicationFlagsEmbedded                      ApplicationFlags = 1 << 17
	ApplicationFlagsGatewayMessageContent         ApplicationFlags = 1 << 18
	ApplicationFlagsGatewayMessageContentLimited  ApplicationFlags = 1 << 19
	ApplicationFlagsApplicationCommandBadge       ApplicationFlags = 1 << 23
)

type InstallParams struct {
	Scopes      []string    `json:"scopes"`
	Permissions Permissions `json:"permissions"`
}

type Team struct {
	Icon        Nullable[string] `json:"icon"`
	ID          Snowflake        `json:"id"`
	Members     []TeamMember     `json:"members"`
	Name        string           `json:"name"`
	OwnerUserID Snowflake        `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState MembershipState    `json:"membership_state"`
	TeamID          Snowflake          `json:"team_id"`
	User            User               `json:"user"`
	Role            TeamMemberRoleType `json:"role"`
}

type MembershipState int

const (
	MembershipStateInvited  MembershipState = 1
	MembershipStateAccepted MembershipState = 2
)

type TeamMemberRoleType string

const (
	TeamMemberRoleTypeAdmin     TeamMemberRoleType = "admin"
	TeamMemberRoleTypeDeveloper TeamMemberRoleType = "developer"
	TeamMemberRoleReadOnly      TeamMemberRoleType = "read_only"
)

type ApplicationCommandPermissions struct{}

type ApplicationCommandPermissionsUpdateEvent = ApplicationCommandPermissions
