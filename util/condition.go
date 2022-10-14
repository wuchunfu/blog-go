package util

import "strings"

func GetCondition(where map[string]any) (query string, args []any) {
	var sb strings.Builder
	for k, v := range where {
		condition := strings.Split(k, " ")
		if sb.String() != "" {
			sb.WriteString(" AND ")
		}
		switch len(condition) {
		case 1:
			sb.WriteString(condition[0])
			sb.WriteString(" = ?")
			args = append(args, v)
			break
		case 2:
			field := condition[0]
			switch condition[1] {
			case "=":
				sb.WriteString(field)
				sb.WriteString(" = ?")
				args = append(args, v)
				break
			case ">":
				sb.WriteString(field)
				sb.WriteString(" > ?")
				args = append(args, v)
				break
			case ">=":
				sb.WriteString(field)
				sb.WriteString(" >= ?")
				args = append(args, v)
				break
			case "<":
				sb.WriteString(field)
				sb.WriteString(" < ?")
				args = append(args, v)
				break
			case "<=":
				sb.WriteString(field)
				sb.WriteString(" <= ?")
				args = append(args, v)
				break
			case "in":
				sb.WriteString(field)
				sb.WriteString(" in ?")
				args = append(args, v)
				break
			case "like":
				sb.WriteString(field)
				sb.WriteString(" like ?")
				args = append(args, v)
				break
			case "<>":
				sb.WriteString(field)
				sb.WriteString(" <> ?")
				args = append(args, v)
				break
			case "!=":
				sb.WriteString(field)
				sb.WriteString(" != ?")
				args = append(args, v)
				break
			}
			break
		}
	}
	query = sb.String()
	return
}
