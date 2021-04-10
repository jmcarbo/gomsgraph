package models

type AutomaticRepliesSetting struct {
	// The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
	ExternalAudience *ExternalAudienceScope `json:"externalAudience,omitempty"`
	// The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
	ExternalReplyMessage *string `json:"externalReplyMessage,omitempty"`
	// The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
	InternalReplyMessage *string `json:"internalReplyMessage,omitempty"`
	// The date and time that automatic replies are set to end, if Status is set to Scheduled.
	ScheduledEndDateTime *DateTimeTimeZone `json:"scheduledEndDateTime,omitempty"`
	// The date and time that automatic replies are set to begin, if Status is set to Scheduled.
	ScheduledStartDateTime *DateTimeTimeZone `json:"scheduledStartDateTime,omitempty"`
	// Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
	Status *AutomaticRepliesStatus `json:"status,omitempty"`
}
