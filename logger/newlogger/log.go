package newlogger

import (
	"github.com/ormushq/ormus/logger/newlogger/loggerparam"
	"github.com/ormushq/ormus/pkg/trace/newtrace"
)

type Log struct {
	cat   loggerparam.Category
	sub   loggerparam.SubCategory
	msg   string
	extra map[loggerparam.ExtraKey]interface{}
}

func NewLog(msg string) *Log {
	return &Log{
		cat:   loggerparam.CategoryNotDefined,
		sub:   loggerparam.SubCategoryNotDefined,
		msg:   msg,
		extra: map[loggerparam.ExtraKey]interface{}{},
	}
}

func (l *Log) WithTrace() *Log {
	return l.With(loggerparam.ExtraKeyTrace, newtrace.Parse(0))
}

func (l *Log) WithCategory(cat loggerparam.Category) *Log {
	l.cat = cat

	return l
}

func (l *Log) WithSubCategory(sub loggerparam.SubCategory) *Log {
	l.sub = sub

	return l
}

func (l *Log) With(key loggerparam.ExtraKey, value interface{}) *Log {
	l.extra[key] = value

	return l
}

func (l *Log) Debug() {
	L().Debug(l.cat, l.sub, l.msg, l.extra)
}

func (l *Log) Info() {
	L().Info(l.cat, l.sub, l.msg, l.extra)
}

func (l *Log) Warn() {
	L().Warn(l.cat, l.sub, l.msg, l.extra)
}

func (l *Log) Error() {
	L().Error(l.cat, l.sub, l.msg, l.extra)
}

func (l *Log) Fatal() {
	L().Fatal(l.cat, l.sub, l.msg, l.extra)
}
