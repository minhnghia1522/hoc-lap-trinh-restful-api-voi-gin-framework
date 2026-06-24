# Lesson 05 - Exercise User Manager

## Flow xử lý

```text
Client -> Router -> Handler -> Validation -> Service -> Repository -> Model
```

### Luồng xử lý

1. Client gửi request đến API.
2. Router nhận request và định tuyến đến handler phù hợp.
3. Handler nhận input, parse `payload` / `query` / `path param` và gọi validation nếu cần.
4. Validation kiểm tra dữ liệu đầu vào.
5. Nếu dữ liệu không hợp lệ, handler trả về response lỗi ngay.
6. Nếu dữ liệu hợp lệ, handler chuyển xử lý nghiệp vụ cho service.
7. Service xử lý logic chính của bài toán, có thể kết hợp nhiều bước nghiệp vụ nếu cần.
8. Service gọi repository để thao tác với dữ liệu.
9. Repository làm việc trực tiếp với storage / database.
10. Repository trả kết quả về service, service có thể chuyển đổi dữ liệu trước khi trả tiếp.
11. Handler build response và trả kết quả về client.

### Vai trò từng layer

- `Router`: khai báo endpoint và gắn với handler.
- `Handler`: tiếp nhận request, validate input, trả response.
- `Validation`: kiểm tra và chuẩn hóa dữ liệu đầu vào.
- `Service`: chứa business logic.
- `Repository`: truy cập dữ liệu.
- `Model`: định nghĩa cấu trúc dữ liệu.

### Validation

Project có thêm layer `validation` để:

- đăng ký custom validator cho `slug`, `search`, `min_int`, `max_int`, `file_ext`.
- chuẩn hóa thông báo lỗi validation thành format dễ đọc hơn.
- giúp handler trả về lỗi rõ ràng khi request không đúng định dạng.

### Tóm tắt

Flow chính của lesson này là request đi từ `router` sang `handler`, sau đó xuống `validation`, `service`, `repository` và cuối cùng thao tác với `model` / dữ liệu.
