package patterns

import "fmt"

type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	kind     string
}

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	Kind     string
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (builder *NotificationBuilder) Build() (*Notification, error) {
	if builder.Icon != "" && builder.Subtitle == "" {
		return nil, fmt.Errorf("icon cannot be used without a subtitle")
	}

	if builder.Priority > 5 || builder.Priority < 0 {
		return nil, fmt.Errorf("priority must be between 0 and 5")
	}

	notification := &Notification{
		title:    builder.Title,
		subtitle: builder.Subtitle,
		message:  builder.Message,
		image:    builder.Image,
		icon:     builder.Icon,
		priority: builder.Priority,
		kind:     builder.Kind,
	}

	return notification, nil
}

func (builder *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	builder.Title = title
	return builder
}

func (builder *NotificationBuilder) SetSubTitle(subTitle string) *NotificationBuilder {
	builder.Subtitle = subTitle
	return builder
}

func (builder *NotificationBuilder) SetMessage(message string) *NotificationBuilder {
	builder.Message = message
	return builder
}

func (builder *NotificationBuilder) SetImage(image string) *NotificationBuilder {
	builder.Image = image
	return builder
}

func (builder *NotificationBuilder) SetIcon(icon string) *NotificationBuilder {
	builder.Icon = icon
	return builder
}

func (builder *NotificationBuilder) SetPriority(priority int) *NotificationBuilder {
	builder.Priority = priority
	return builder
}

func (builder *NotificationBuilder) SetKind(kind string) *NotificationBuilder {
	builder.Kind = kind
	return builder
}
