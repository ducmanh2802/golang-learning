# Project's Info
|     Fullname    |    Class     |
|-----------------|--------------|
| Dang Duc Manh  |   TFS - 03   |
Golang basic - Kiến thức cơ bản về golang
- là ngôn ngữ lập trình biên dịch
- IDE: VS Code/ Goland

Một số lệnh git cơ bản:
    git init
    git status
    git commit
    git push

Concurrence vs Parallelism:
  - Concurrence: tận dụng khoảng gián đoạn của các task để nhét thêm task khác vào thực thi(Go thường sử dụng để tránh tốn process)
  - Parallelism: chạy song song, sử dụng nhiều process
  => thời gian của cái nào nhanh hơn?? => tùy dự án

Types
  - bool, string,int int8 int 16 int 32 int 64,byte, float32 float64,
  Cách khai báo:
  i:=1
  var i, j int = 1, 2

  Collection Types
    - Array số lượng phần tử cố định
    var arr = [2] int {}
    var arr = [2] int {1,2}
    var arr = [...] int {1,2,3,4} //Có thể viết thế này để khi khởi tạo ko quan tâm tới số lượng phần tử
    Pass by reference(Slice): truyền biến đó vào 1 hàm thì chỉ truyền giá trị
    Pass by value(Array) clone ra 1 giá trị mới, giá trị ngoài hàm ko thay đổi => tốn bộ nhớ, chi phí hơn, do đó slice hay được sử dụng hơn

    Control flow(if, for, switch)
      - bỏ dấu ngoặc (), logic vẫn giống java
      - Gollang ko có While
      - for ko điều kiện sẽ chạy vòng lặp vô hạn


1. Naming Conventions in Go: Short but Descriptive
