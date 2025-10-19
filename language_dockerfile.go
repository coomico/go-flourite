package flourite

import "strings"

var dockerKeywords = []string{
	"ADD",
	"ARG",
	"AS",
	"CMD",
	"COPY",
	"CROSS_BUILD",
	"ENTRYPOINT",
	"ENV",
	"EXPOSE",
	"FROM",
	"HEALTHCHECK",
	"LABEL",
	"MAINTAINER",
	"ONBUILD",
	"RUN",
	"SHELL",
	"STOPSIGNAL",
	"USER",
	"VOLUME",
	"WORKDIR",
}

var dockerExpression = "^(" + strings.Join(dockerKeywords, "|") + ")"

var dockerfile = []languagePattern{
	{expression: dockerExpression, patternType: keyword},
}
