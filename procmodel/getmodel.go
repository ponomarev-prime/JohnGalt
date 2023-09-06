package procmodel

import (
	"strings"
)

func GetModelName(cpuInfo string) string {
	lines := strings.Split(cpuInfo, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "Информация о модели процессора не найдена"
}
