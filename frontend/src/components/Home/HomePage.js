import React, { useState, Fragment } from "react";
import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";
import Notifications from "../Notification/Notifications";
import { isExpired } from "react-jwt";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import "./Home.css";
import HeaderBar from "../Navbar/HeaderBar"

const HomePage = () => {
  const [date, setDate] = useState(new Date());
  const history = useHistory();
  const token = localStorage.getItem("token");

  if (!token || isExpired(token)) history.push("/login");

  return (
    <Fragment>
      <HeaderBar />
      <div className="section">
        <Calendar onChange={(date) => setDate(date)} value={date} />
        <Notifications date={date} />
      </div>
    </Fragment>
  );
};

export default HomePage
