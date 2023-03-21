package schema

// SharepointBase represents the base schema for sharepoint records.
type SharepointBase struct {
	AuditRecord
	Site                   *string `json:"Site,omitempty"`
	ItemType               *string `json:"ItemType,omitempty"`
	EventSource            *string `json:"EventSource,omitempty"`
	SourceName             *string `json:"SourceName,omitempty"`
	UserAgent              *string `json:"UserAgent,omitempty"`
	MachineDomainInfo      *string `json:"MachineDomainInfo,omitempty"`
	MachineID              *string `json:"MachineId,omitempty"`
	ListItemUniqueID       *string `json:"ListItemUniqueId,omitempty"`
	ListID                 *string `json:"ListId,omitempty"`
	ApplicationID          *string `json:"ApplicationId,omitempty"`
	ApplicationDisplayName *string `json:"ApplicationDisplayName,omitempty"`
	IsWorkflow             *bool   `json:"IsWorkflow"`
}

// SharepointFileOperations .
type SharepointFileOperations struct {
	SharepointBase
	SiteURL                  *string `json:"SiteUrl"`
	SourceRelativeURL        *string `json:"SourceRelativeUrl,omitempty"`
	SourceFileName           *string `json:"SourceFileName"`
	SourceFileExtension      *string `json:"SourceFileExtension,omitempty"`
	DestinationRelativeURL   *string `json:"DestinationRelativeUrl,omitempty"`
	DestinationFileName      *string `json:"DestinationFileName,omitempty"`
	DestinationFileExtension *string `json:"DestinationFileExtension,omitempty"`
	UserSharedWith           *string `json:"UserSharedWith,omitempty"`
	SharingType              *string `json:"SharingType,omitempty"`
}

// SharepointSharing .
type SharepointSharing struct {
	SharepointBase
	TargetUserOrGroupName *string `json:"TargetUserOrGroupName,omitempty"`
	TargetUserOrGroupType *string `json:"TargetUserOrGroupType,omitempty"`
	EventData             *string `json:"EventData,omitempty"`
}

// Sharepoint .
type Sharepoint struct {
	SharepointBase
	CustomEvent        *string  `json:"CustomEvent,omitempty"`
	EventData          *string  `json:"EventData,omitempty"`
	ModifiedProperties []string `json:"ModifiedProperties,omitempty"`
}
