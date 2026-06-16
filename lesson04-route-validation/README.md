📌 Nội dung chính của bài học này:
✅ Tại sao Router Validation lại quan trọng? (Ngăn chặn hacker, đảm bảo dữ liệu đầu vào hợp lệ).
✅ Trường hợp 1: Validate Path Parameter là số nguyên dương.
✅ Lấy Path Param, chuyển đổi sang số (strconv.Atoi).
✅ Kiểm tra lỗi chuyển đổi (có phải là số không?).
✅ Kiểm tra giá trị có lớn hơn 0 không.
✅ Trả về lỗi 400 Bad Request nếu không hợp lệ.
✅ Trường hợp 2: Validate Path Parameter là UUID (Universally Unique Identifier).
✅ Sử dụng package github.com/google/uuid để parse và validate UUID.
✅ Cài đặt package: go get github.com/google/uuid.
✅ Trả về lỗi nếu chuỗi đầu vào không phải là một UUID hợp lệ.
✅ Trường hợp 3: Validate Path Parameter là Slug.
✅ Slug là gì? (Chuỗi thân thiện với URL, thường dùng cho tiêu đề bài viết, tên sản phẩm).
✅ Quy tắc cho slug: chỉ chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm, không có ký tự đặc biệt, không có nhiều dấu gạch ngang/chấm liên tiếp.
✅ Sử dụng Biểu thức chính quy (Regular Expression - Regex) để validate slug.
✅ Package regexp trong Golang
✅ Trường hợp 4: Validate Path Parameter là một giá trị cụ thể trong danh sách cho phép (Enum-like validation).
✅ Ví dụ: Category chỉ được phép là "php", "python", hoặc "golang".
✅ Sử dụng map để định nghĩa các giá trị hợp lệ và kiểm tra.
✅ Trường hợp 5: Validate Query Parameters.
✅ Ví dụ: search term và limit.
✅ Validate search: không được rỗng, độ dài ký tự trong khoảng cho phép, chỉ chứa chữ cái, số và khoảng trắng (dùng Regex).
✅ Validate limit: nếu không truyền thì có giá trị mặc định (sử dụng c.DefaultQuery), phải là số nguyên dương.
✅ Cách trả về các thông báo lỗi rõ ràng cho client khi validation thất bại.
✅ Thực hành trực tiếp trên project đã có, áp dụng validation cho các router hiện tại.

Router Validation là một lớp phòng thủ quan trọng cho API của bạn. Bằng cách triển khai các quy tắc kiểm tra chặt chẽ, bạn không chỉ nâng cao chất lượng API mà còn bảo vệ hệ thống của mình khỏi các rủi ro tiềm ẩn.