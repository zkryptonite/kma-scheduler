import React, { useState, useEffect } from "react";
import Notification from "./Notification";


const Notifications = ({ date }) => {
    const [notifications, setNotifications] = useState([]);
    const formattedDay = `${date.getDate()}/${date.getMonth() + 1}/${date.getFullYear()}`;
    const headers = new Headers();
    const fetchUrl = "http://localhost:5000/api/classes";
    headers.append(
        "Authorization",
        `Bearer ${localStorage.getItem("token")}`
    );

    const requestOptions = {
        method: "GET",
        headers: headers,
        mode: "cors",
        redirect: "follow",
    };

    useEffect(() => {
        (async () => {
            const res = await fetch(`${fetchUrl}?date=${formattedDay}`, requestOptions);
            setNotifications(await res.json());
        })();
    }, [date]);
    
    if (notifications === null) {
        return (
            <div className="notification-container">
                <div className="no-notification">
                    Không có lịch học
                </div>
            </div>
        )
    }


    return (
        <div className="notification-container">
            {notifications.map(notification => (
                <Notification notificationProps={notification} key={notification.subject_name}/>
            ))}
        </div>
    )
}

export default Notifications