package templates

import "time"

func TimeAgo(t time.Time) string {
	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "ahora mismo"
	case d < time.Hour:
		m := int(d.Minutes())
		if m == 1 {
			return "hace 1 minuto"
		}
		return "hace " + itoa(m) + " minutos"
	case d < 24*time.Hour:
		h := int(d.Hours())
		if h == 1 {
			return "hace 1 hora"
		}
		return "hace " + itoa(h) + " horas"
	case d < 7*24*time.Hour:
		dias := int(d.Hours() / 24)
		if dias == 1 {
			return "hace 1 día"
		}
		return "hace " + itoa(dias) + " días"
	case d < 30*24*time.Hour:
		semanas := int(d.Hours() / (24 * 7))
		if semanas == 1 {
			return "hace 1 semana"
		}
		return "hace " + itoa(semanas) + " semanas"
	case d < 365*24*time.Hour:
		meses := int(d.Hours() / (24 * 30))
		if meses == 1 {
			return "hace 1 mes"
		}
		return "hace " + itoa(meses) + " meses"
	default:
		anos := int(d.Hours() / (24 * 365))
		if anos == 1 {
			return "hace 1 año"
		}
		return "hace " + itoa(anos) + " años"
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	i := len(buf)
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
