# PROJECT KMA SCHEDULER

## 1. Giới Thiệu
Xuất phát từ thực tế trang web dành cho sinh viên của trường khó theo dõi lịch học, giao diện khá cũ và có ít tính năng bổ sung, để khắc phục những hạn chế tôi đã xây dựng nên một ứng dụng web nhỏ. Ứng dụng web sử dụng dữ liệu lịch học của chính sinh viên trường nên nó có thể được ứng dụng vào thực tế ngay lập tức, có tính thực tiễn cao với hi vọng có thể giúp sinh viên có thể sắp xếp, theo dõi lịch học một cách dễ dàng hỗ trợ cho việc sắp xếp thời gian biểu hằng ngày một cách hợp lí hơn, từ đó để có một ngày làm việc năng suất, hiệu quả hơn.

## 2. Các Chức Năng
* Trang đăng nhập sử dụng chính tài khoản sinh viên của trường nên không cần thêm các bước setup phức tạp. Ứng dụng sẽ tự động crawl dữ liệu lịch học của sinh viên về lưu vào database sau khi đăng nhập thành công.
   
![kma-scheduler](/docs/A4370117-009D-4847-9F51-18B0EC3D004E.png)
* Trang chính để theo dõi lịch học, ngoài ra sinh viên cũng có thể đăng kí email để nhận thông báo về lịch học hàng ngày.
   
![kma-scheduler](/docs/5300CE2B-44D5-4E15-9571-C494A4FDD0C4.png)
* Nhận thông báo lịch học hàng ngày khi đăng ký email.

![kma-scheduler](/docs/56DE9263-66AB-4240-BF25-450736F85ABF.png)


## 3. Run Project:
##### Để chạy project thực hiện các bước sau:

     1. Clone project về máy local.

     2. Navigate tới project trên terminal chạy câu lệnh "docker-compose build" 
        để build project thành docker image.

     3. Sau đó sử dụng câu lệnh "docker-compose run" để khởi động và chạy container.
 
     4. Ứng dụng được chạy trên trình duyệt với port 3000.