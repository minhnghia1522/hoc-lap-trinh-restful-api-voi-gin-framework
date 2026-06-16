📌 Nội dung chính của bài học này:
✅ Giới thiệu khái niệm Router Group trong Gin Framework.
✅ Tại sao cần Router Group? Lợi ích của việc nhóm các routes.
✅ Ví dụ thực tế: Quản lý API cho user và product với các phiên bản (version) khác nhau (V1, V2).
✅ Bước 1: Thiết lập project, cài đặt Gin và tạo server cơ bản.
✅ Bước 2: Tổ chức code với cấu trúc thư mục khoa học:
✅ Tạo thư mục internal/api để chứa logic API.
✅ Phân chia theo version (ví dụ: v1, v2).
✅ Trong mỗi version, tạo thư mục handler để chứa các file xử lý cho từng resource (ví dụ: user.go, product.go).
✅ Bước 3: Viết các handler functions cho User và Product (GET all, GET by ID, POST, PUT, DELETE).
✅ Tách logic handler ra các file và package riêng biệt.
✅ Sử dụng struct và method receivers cho handler để tổ chức code tốt hơn.
✅ Bước 4: Đăng ký routes mà không sử dụng Router Group (cách làm dài dòng, lặp lại tiền tố).
✅ Bước 5: Giới thiệu và áp dụng Router Group:
✅ Sử dụng router.Group("/api/v1") để tạo một group cho version 1.
✅ Các routes trong group sẽ tự động có tiền tố /api/v1.
✅ Cách đăng ký routes bên trong một group (ví dụ: v1.GET("/users", ...)).
✅ Bước 6: Áp dụng Router Group lồng nhau (Nested Router Groups):
✅ Ví dụ: v1UserGroup := v1.Group("/users"), sau đó v1UserGroup.GET("", ...) và v1UserGroup.GET("/:id", ...).
✅ Giúp code trở nên ngắn gọn và dễ quản lý hơn nữa.
✅ Thực hành refactor code để sử dụng Router Group cho tất cả các API (User V1, Product V1, User V2).
✅ So sánh code trước và sau khi sử dụng Router Group: thấy rõ sự cải thiện về tính rõ ràng và khả năng bảo trì.