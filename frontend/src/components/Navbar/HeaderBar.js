import React, { useState } from "react";
import "./HeaderBar.css"
import { Navbar, Nav, Container, FormControl, Form, Button } from "react-bootstrap"
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSignOutAlt, faCalendarAlt } from '@fortawesome/fontawesome-free-solid'

const HeaderBar = () => {
  const history = useHistory();
  const [email, setEmail] = useState("");

  const logout = () => {
    localStorage.clear();
    history.push("/login");
  }

  const subcribe = async (event) => {
    event.preventDefault();
    let header = new Headers();
    header.append("Content-Type", "application/json");
    header.append("Authorization", `Bearer ${localStorage.getItem("token")}`);

    let requestOptions = {
      method: 'POST',
      headers: header,
      body: JSON.stringify({email: email}),
      redirect: 'follow'
    };

    await fetch("http://localhost:5000/api/student/email", requestOptions)
      .then(response => {
        if (response.status === 200) alert("Đăng email thành công!");
        else alert("Đăng kí thất bại!"); 
      })
      .catch(error => console.log('error', error));
  }


  return (
    <Navbar className="header-bar" bg="light" variant="light" expand="md">
      <Container fluid>
        <Navbar.Brand href="#"><FontAwesomeIcon size="1x" icon={faCalendarAlt} /> Thời Khoá Biểu</Navbar.Brand>
        <Navbar.Toggle aria-controls="navbarScroll" />
        <Navbar.Collapse id="navbarScroll">
          <Nav
            className="me-auto my-2 my-lg-0"
            style={{ maxHeight: "100px" }}
            navbarScroll
          >
            <Nav.Link onClick={logout}>LogOut  <FontAwesomeIcon icon={faSignOutAlt} /></Nav.Link>
          </Nav>
          <Form className="d-flex" onSubmit={subcribe}>
            <FormControl
              type="email"
              placeholder="Đăng kí email"
              onChange={event => setEmail(event.target.value)}
              className="me-2"
            />
            <Button type="submit" variant="primary">Subcribe</Button>
          </Form>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

export default HeaderBar