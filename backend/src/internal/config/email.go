package config

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

// Отправка писем через SMTP
type SMTPEmailService struct {
	host     string
	port     string
	username string
	password string
	from     string
}

// Создание нового экземпляра SMTPEmailService
func NewSMTPEmailService(host, port, username, password, from string) *SMTPEmailService {
	return &SMTPEmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

// Создает и настраивает SMTP сервис из env
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

// Интерфейс для отправки email
type EmailSender interface {
	SendVerificationCode(toEmail, code string) error
	SendVerificationResetCode(toEmail, code string) error
}

// Отправление кода подтверждения для регистрации
func (s *SMTPEmailService) SendVerificationCode(toEmail, code string) error {
	subject := "Код подтверждения регистрации"

	// Тело письма
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
  </head>

  <body
    style="
      margin: 0;
      padding: 0;
      background: #fffdf7;
      font-family: Arial, sans-serif;
      color: #2c3e50;
    "
  >
    <table width="100%" cellpadding="0" cellspacing="0" style="padding: 40px 0">
      <tr>
        <td align="center">
          <table
            width="520"
            cellpadding="0"
            cellspacing="0"
            style="
              background: #ffffff;
              border-radius: 32px;
              overflow: hidden;
              border: 1px solid rgba(88, 204, 2, 0.4);
              box-shadow: 0 12px 30px rgba(88, 204, 2, 0.08);
            "
          >
            <tr>
              <td style="padding: 28px 30px; text-align: center; background: #58cc02">
                <div
                  style="
                    font-size: 28px;
                    font-weight: bold;
                    letter-spacing: 2px;
                    color: white;
                  "
                >
                  TeachLangua
                </div>
              </td>
            </tr>

            <tr>
              <td style="padding: 35px 40px; text-align: left">
                <h2 style="margin-top: 0; color: #2c3e50">Код подтверждения</h2>

                <p style="color: #555; line-height: 1.6">
                  Чтобы завершить регистрацию, используйте следующий код подтверждения:
                </p>

                <div
                  style="
                    font-size: 40px;
                    font-weight: bold;
                    letter-spacing: 8px;
                    margin: 30px 0;
                    padding: 18px 25px;
                    background: #f9fff0;
                    border: 2px dashed #58cc02;
                    border-radius: 20px;
                    text-align: center;
                    color: #2c3e50;
                  "
                >
                  %s
                </div>

                <p style="color: #777">
                  Код действителен в течение <b>10 минут</b>.
                </p>

                <p style="color: #777; margin-top: 30px">
                  С уважением,<br />
                  <b style="color: #2c3e50">Команда TeachLangua</b>
                </p>
              </td>
            </tr>

            <tr>
              <td
                style="
                  padding: 18px;
                  text-align: center;
                  font-size: 12px;
                  color: #999;
                  background: #fcfcf5;
                  border-top: 1px solid rgba(88, 204, 2, 0.2);
                "
              >
                Если вы не регистрировались, просто проигнорируйте это письмо.
              </td>
            </tr>

            <tr>
              <td
                style="
                  padding: 16px;
                  text-align: center;
                  font-size: 12px;
                  color: #aaa;
                "
              >
                © %d TeachLangua
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`, code, time.Now().Year())

	// Заголовки
	headers := make(map[string]string)
	headers["From"] = s.from
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Сборка сообщения
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Аутентификация
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	// Адрес сервера
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	// Отправление с таймаутом
	errChan := make(chan error, 1)
	go func() {
		errChan <- smtp.SendMail(addr, auth, s.from, []string{toEmail}, []byte(message))
	}()

	// Ожидание ответа (15 секунд)
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

// Отправление кода подтверждения для восстановления пароля
func (s *SMTPEmailService) SendVerificationResetCode(toEmail, code string) error {
	subject := "Код подтверждения для восстановления пароля"

	// Тело письма
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
  </head>

  <body style="margin:0; padding:0; background:#0f1c1e; font-family:Arial, sans-serif; color:white;">

    <table width="100%%" cellpadding="0" cellspacing="0" style="padding:40px 0;">
      <tr>
        <td align="center">

          <table width="520" cellpadding="0" cellspacing="0" style="background:#1c2a2c; border-radius:16px; overflow:hidden;">

            <tr>
              <td style="padding:30px; text-align:center; background:rgb(22,71,71);">
                <div style="font-size:26px; font-weight:bold; letter-spacing:2px;">
                  PIONEER
                </div>
              </td>
            </tr>

            <tr>
              <td style="padding:35px 40px; text-align:left;">

                <h2 style="margin-top:0; color:white;">
                  Восстановление пароля
                </h2>

                <p style="color:#cfd8dc; line-height:1.6;">
                  Для восстановления пароля используйте следующий код подтверждения:
                </p>

                <div style="
                  font-size:40px;
                  font-weight:bold;
                  letter-spacing:8px;
                  margin:30px 0;
                  padding:18px 25px;
                  background:rgb(22,71,71);
                  border-radius:10px;
                  text-align:center;
                ">
                  %s
                </div>

                <p style="color:#cfd8dc;">
                  Код действителен в течение <b>10 минут</b>.
                </p>

                <p style="color:#cfd8dc; margin-top:30px;">
                  С уважением,<br>
                  <b>Pioneer</b>
                </p>

              </td>
            </tr>

            <tr>
              <td style="padding:20px; text-align:center; font-size:12px; color:#9ea7aa; background:#182426;">
                Если вы не запрашивали восстановление пароля, просто проигнорируйте это письмо.
              </td>
            </tr>

            <tr>
              <td style="padding:20px; text-align:center; font-size:12px; color:#9ea7aa;">
                © %d Pioneer
              </td>
            </tr>

          </table>

        </td>
      </tr>
    </table>

  </body>
</html>
`, code, time.Now().Year())

	// Заголовки
	headers := make(map[string]string)
	headers["From"] = s.from
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Сборка сообщения
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Аутентификация
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	// Адрес сервера
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	// Отправление с таймаутом
	errChan := make(chan error, 1)
	go func() {
		errChan <- smtp.SendMail(addr, auth, s.from, []string{toEmail}, []byte(message))
	}()

	// Ожидание ответа (15 секунд)
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
