import React from "react";
import "./Notification.css";

const Notification = ({ notificationProps }) => {
  return (
    <div className="notification-item">
      Môn học: {notificationProps.subject_name}
      <br />
      Giảng viên:{" "}
      {notificationProps.lecturer === ""
        ? "Không có thông tin"
        : notificationProps.lecturer}
      <br />
      Thời gian: {notificationProps.lesson}
      <br />
      Phòng học:{" "}
      {notificationProps.room === ""
        ? "Không có thông tin"
        : notificationProps.room}
    </div>
  );
};

export default Notification;
