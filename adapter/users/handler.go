package UsersAdapter

import (
	UsersModels "headliner-be/model/users"
	UsersUsecase "headliner-be/usecase/users"
	"headliner-be/utils"

	fbauth "firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v3"
)

type UsersHandler struct {
    UsersUsecase UsersUsecase.UsersUsecase
}

func NewUsersAdapter(usersUsecase UsersUsecase.UsersUsecase) *UsersHandler {
    return &UsersHandler{
        UsersUsecase: usersUsecase,
    }
}

// Register godoc
// @Summary      Register
// @Description  ใช้สำหรับลงทะเบียนผู้ใช้ใหม่เข้าสู่ระบบ
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.RegisterInput true "ข้อมูลการสมัครสมาชิก"
// @Success      200  {object}  map[string]interface{} "สมัครสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/register [post]
func (a *UsersHandler) Register(c fiber.Ctx) error {
    var users UsersModels.RegisterInput
    if err := c.Bind().Body(&users); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }
    if !utils.IsValidEmail(users.Email) {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid email format", "", nil)
    }
    message, err := a.UsersUsecase.Register(&users)
    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, message, err.Error(), nil)
    }
    return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}

// Login godoc
// @Summary      Login
// @Description  ใช้สำหรับล็อกอินเข้าเกมและรับ Token
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.LoginInput true "ข้อมูลการเข้าสู่ระบบ"
// @Success      200  {object}  map[string]interface{} "ล็อกอินสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/login [post]
func (a *UsersHandler) Login(c fiber.Ctx) error {
    var users UsersModels.LoginInput
    if err := c.Bind().Body(&users); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }
    data, err := a.UsersUsecase.Login(&users)
    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, data, err.Error(), nil)
    }
    return utils.ResponseJSON(c, fiber.StatusOK, "Login successful", "", fiber.Map{
        "token": data,
    })
}

// CreateCharacter godoc
// @Summary      สร้างตัวละคร
// @Description  ใช้สำหรับสร้างหรือตั้งค่าตัวละครของผู้เล่น
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.CreateCharacterInput true "ข้อมูลตัวละคร"
// @Success      200  {object}  map[string]interface{} "สร้างตัวละครสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/createcharacter [post]
func (a *UsersHandler) CreateCharacter(c fiber.Ctx) error {
    var users UsersModels.CreateCharacterInput
    if err := c.Bind().Body(&users); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }
    userID := c.Locals("userID")

    message, err := a.UsersUsecase.CreateCharacter(uint(userID.(float64)), &users)
    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, message, err.Error(), nil)
    }
    return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}

// GetUserData godoc
// @Summary      ดึงข้อมูลผู้ใช้งาน
// @Description  ดึงข้อมูลโปรไฟล์ ตัวละคร และชื่อบอทของผู้เล่นปัจจุบัน
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/data [get]
func (a *UsersHandler) GetUserData(c fiber.Ctx) error {
    userID := c.Locals("userID")

    user, err := a.UsersUsecase.GetUserData(uint(userID.(float64)))
    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get user data", err.Error(), nil)
    }
    return utils.ResponseJSON(c, fiber.StatusOK, "User data retrieved successfully", "", fiber.Map{
        "id":           user.ID,
        "username":     user.Username,
        "email":        user.Email,
        "character":    user.Character,
        "chatbot_name": user.Chatbotname, // เผื่อส่งกลับไปให้ Unity ใช้ตอน Login
    })
}

// UpdateUsername godoc
// @Summary      อัปเดตชื่อผู้ใช้งาน
// @Description  เปลี่ยนชื่อผู้ใช้งาน (Username) ของผู้เล่น
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.UpdateUsernameInput true "ชื่อผู้ใช้งานใหม่"
// @Success      200  {object}  map[string]interface{} "อัปเดตสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Router       /users/update-username [put]
func (a *UsersHandler) UpdateUsername(c fiber.Ctx) error {
    var input UsersModels.UpdateUsernameInput

    if err := c.Bind().Body(&input); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }

    userID := c.Locals("userID")

    message, err := a.UsersUsecase.UpdateUsername(uint(userID.(float64)), &input)

    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, message, err.Error(), nil)
    }

    return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}

// UpdatePassword godoc
// @Summary      อัปเดตรหัสผ่าน
// @Description  เปลี่ยนรหัสผ่านของผู้เล่น
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.UpdatePasswordInput true "รหัสผ่านใหม่"
// @Success      200  {object}  map[string]interface{} "อัปเดตสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Router       /users/update-password [put]
func (a *UsersHandler) UpdatePassword(c fiber.Ctx) error {
    var input UsersModels.UpdatePasswordInput

    if err := c.Bind().Body(&input); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }

    userID := c.Locals("userID")

    message, err := a.UsersUsecase.UpdatePassword(uint(userID.(float64)), &input)

    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, message, err.Error(), nil)
    }

    return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}

// UpdateChatbotName godoc
// @Summary      Update Chatbot Name
// @Description  เปลี่ยนชื่อแชทบอทที่ปรึกษาทางการเงินของผู้เล่น
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body UsersModels.UpdateChatbotRequest true "ชื่อแชทบอทใหม่"
// @Success      200  {object}  map[string]interface{} "อัปเดตสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/update-chatbot [post]
func (a *UsersHandler) UpdateChatbotName(c fiber.Ctx) error {
    var input UsersModels.UpdateChatbotRequest

    if err := c.Bind().Body(&input); err != nil {
        return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
    }

    userID := c.Locals("userID")

    err := a.UsersUsecase.UpdateChatbotName(uint(userID.(float64)), input.ChatbotName)

    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update chatbot name", err.Error(), nil)
    }

    return utils.ResponseJSON(c, fiber.StatusOK, "Chatbot name updated successfully", "", fiber.Map{
        "chatbot_name": input.ChatbotName,
    })
}

// LoginWithGoogle godoc
// @Summary      Login with Google
// @Description  ล็อกอินเข้าเกมผ่าน Firebase Google Auth
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]interface{} "ล็อกอินสำเร็จ"
// @Failure      401  {object}  map[string]interface{} "ไม่มีสิทธิ์เข้าถึง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /users/login-with-google [post]
func (a *UsersHandler) LoginWithGoogle(c fiber.Ctx) error {
    // 1. ดึงข้อมูลจาก Locals (ต้องสะกด Key ให้ตรงกับใน Middleware)
    rawToken := c.Locals("GOOGLE_AUTH_TOKEN") 
    
    // 2. ตรวจสอบว่ามีข้อมูลไหม (ป้องกัน Nil Pointer)
    if rawToken == nil {
        return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Unauthorized", "Firebase token not found in context", nil)
    }

    // 3. Type Assertion แปลง interface{} เป็น *auth.Token
    googleAuthToken, ok := rawToken.(*fbauth.Token)
    if !ok {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Internal Server Error", "Invalid token type", nil)
    }

    // 4. ส่งเข้าไปใน Usecase
    token, err := a.UsersUsecase.LoginWithGoogle(googleAuthToken)

    if err != nil {
        return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to login with Google", err.Error(), nil)
    }

    return utils.ResponseJSON(c, fiber.StatusOK, "Login with Google successful", "", fiber.Map{
        "token": token,
    })
}

// RegisterWithGoogle godoc
// @Summary      Register with Google
// @Description  ลงทะเบียนผู้ใช้ใหม่ผ่าน Firebase Google Auth
// @Tags         Users
// @Produce      json
// @Success      201  {object}  map[string]interface{} "สมัครสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      401  {object}  map[string]interface{} "ไม่มีสิทธิ์เข้าถึง"
// @Router       /users/register-with-google [post]
func (a *UsersHandler) RegisterWithGoogle(c fiber.Ctx) error {
    // 1. ดึง Firebase Token จาก Middleware (ใช้ Key เดียวกับที่ตั้งไว้)
    rawToken := c.Locals("GOOGLE_AUTH_TOKEN")
    if rawToken == nil {
        return utils.ResponseJSON(c, 401, "Unauthorized", "Firebase token missing", nil)
    }
    fbUser := rawToken.(*fbauth.Token)

    // 3. เรียก Usecase
    token, err := a.UsersUsecase.RegisterWithGoogle(fbUser)
    if err != nil {
        return utils.ResponseJSON(c, 400, "Registration failed", err.Error(), nil)
    }

    return utils.ResponseJSON(c, 201, "Registered successfully", "", fiber.Map{
        "token": token,
    })
}