package config

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

// SMTPEmailService – отправка писем через SMTP
type SMTPEmailService struct {
	host     string
	port     string
	username string
	password string
	from     string
}

// NewSMTPEmailService создаёт новый экземпляр SMTPEmailService
func NewSMTPEmailService(host, port, username, password, from string) *SMTPEmailService {
	return &SMTPEmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

// ConnectSMTP создаёт и настраивает SMTP сервис из env
func ConnectSMTP() (*SMTPEmailService, error) {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")

	return &SMTPEmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}, nil
}

// EmailSender – интерфейс для отправки email
type EmailSender interface {
	SendVerificationCode(toEmail, code string) error
	SendVerificationResetCode(toEmail, code string) error
}

// SendVerificationCode отправляет код подтверждения регистрации
func (s *SMTPEmailService) SendVerificationCode(toEmail, code string) error {
	subject := "Код подтверждения регистрации"

	// Шаблон письма (светлая тема, TeachLangua)
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
  </head>
  <body style="margin:0;padding:0;background:#fffdf7;font-family:Arial,sans-serif;color:#2c3e50;">
    <table width="100%%" cellpadding="0" cellspacing="0" style="padding:40px 0;">
      <tr>
        <td align="center">
          <table width="520" cellpadding="0" cellspacing="0" style="background:#ffffff;border-radius:32px;overflow:hidden;border:1px solid rgba(88,204,2,0.4);box-shadow:0 12px 30px rgba(88,204,2,0.08);">
            <tr>
              <td style="padding:28px 30px;text-align:center;background:#58cc02;">
                <div style="font-size:28px;font-weight:bold;letter-spacing:2px;color:white;">TeachLangua</div>
              </td>
            </tr>
            <tr>
              <td style="padding:35px 40px;text-align:left;">
                <h2 style="margin-top:0;color:#2c3e50;">Код подтверждения</h2>
                <p style="color:#555;line-height:1.6;">Чтобы завершить регистрацию, используйте следующий код подтверждения:</p>
                <div style="font-size:40px;font-weight:bold;letter-spacing:8px;margin:30px 0;padding:18px 25px;background:#f9fff0;border:2px dashed #58cc02;border-radius:20px;text-align:center;color:#2c3e50;">%s</div>
                <p style="color:#777;">Код действителен в течение <b>10 минут</b>.</p>
                <p style="color:#777;margin-top:30px;">С уважением,<br/><b style="color:#2c3e50;">Команда TeachLangua</b></p>
              </td>
            </tr>
            <tr>
              <td style="padding:18px;text-align:center;font-size:12px;color:#999;background:#fcfcf5;border-top:1px solid rgba(88,204,2,0.2);">Если вы не регистрировались, просто проигнорируйте это письмо.</td>
            </tr>
            <tr>
              <td style="padding:16px;text-align:center;font-size:12px;color:#aaa;">© %d TeachLangua</td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`, code, time.Now().Year()) // Сначала код, потом год

	headers := map[string]string{
		"From":         s.from,
		"To":           toEmail,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	errChan := make(chan error, 1)
	go func() {
		errChan <- smtp.SendMail(addr, auth, s.from, []string{toEmail}, []byte(message))
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("ошибка отправки email: %w", err)
		}
		return nil
	case <-time.After(15 * time.Second):
		return fmt.Errorf("таймаут отправки email")
	}
}

// SendVerificationResetCode отправляет код для восстановления пароля
func (s *SMTPEmailService) SendVerificationResetCode(toEmail, code string) error {
	subject := "Код подтверждения для восстановления пароля"

	// Единый стиль письма (светлая тема, TeachLangua)
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
  </head>
  <body style="margin:0;padding:0;background:#fffdf7;font-family:Arial,sans-serif;color:#2c3e50;">
    <table width="100%%" cellpadding="0" cellspacing="0" style="padding:40px 0;">
      <tr>
        <td align="center">
          <table width="520" cellpadding="0" cellspacing="0" style="background:#ffffff;border-radius:32px;overflow:hidden;border:1px solid rgba(88,204,2,0.4);box-shadow:0 12px 30px rgba(88,204,2,0.08);">
            <tr>
              <td style="padding:28px 30px;text-align:center;background:#58cc02;">
                <div style="font-size:28px;font-weight:bold;letter-spacing:2px;color:white;">TeachLangua</div>
              </td>
            </tr>
            <tr>
              <td style="padding:35px 40px;text-align:left;">
                <h2 style="margin-top:0;color:#2c3e50;">Восстановление пароля</h2>
                <p style="color:#555;line-height:1.6;">Для восстановления пароля используйте следующий код подтверждения:</p>
                <div style="font-size:40px;font-weight:bold;letter-spacing:8px;margin:30px 0;padding:18px 25px;background:#f9fff0;border:2px dashed #58cc02;border-radius:20px;text-align:center;color:#2c3e50;">%s</div>
                <p style="color:#777;">Код действителен в течение <b>10 минут</b>.</p>
                <p style="color:#777;margin-top:30px;">С уважением,<br/><b style="color:#2c3e50;">Команда TeachLangua</b></p>
              </td>
            </tr>
            <tr>
              <td style="padding:18px;text-align:center;font-size:12px;color:#999;background:#fcfcf5;border-top:1px solid rgba(88,204,2,0.2);">Если вы не запрашивали восстановление пароля, просто проигнорируйте это письмо.</td>
            </tr>
            <tr>
              <td style="padding:16px;text-align:center;font-size:12px;color:#aaa;">© %d TeachLangua</td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`, code, time.Now().Year()) // Сначала код, потом год

	headers := map[string]string{
		"From":         s.from,
		"To":           toEmail,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	errChan := make(chan error, 1)
	go func() {
		errChan <- smtp.SendMail(addr, auth, s.from, []string{toEmail}, []byte(message))
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("ошибка отправки email: %w", err)
		}
		return nil
	case <-time.After(15 * time.Second):
		return fmt.Errorf("таймаут отправки email")
	}
}
