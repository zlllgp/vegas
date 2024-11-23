package entity

type ActivityOption func(activity *Activity)

func WithName(name string) ActivityOption {
	return func(activity *Activity) {
		activity.Name = name
	}
}
