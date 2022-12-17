import React, { Fragment, useState } from "react";
import { Container, Row, Col, Form, Button, Alert } from "react-bootstrap";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
import "./LoginForm.css";
import "bootstrap/dist/css/bootstrap.min.css";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUserCircle } from '@fortawesome/fontawesome-free-solid'

const LoginForm = () => {
    const [user, setUser] = useState({ username: "", password: "" });
    const [failed, setFailed] = useState(false);
    const history = useHistory();

    const onSubmitHandler = async () => {
        let header = new Headers();
        header.append("Content-Type", "application/json");

        await fetch("http://localhost:5000/login", {
            method: "POST",
            headers: header,
            body: JSON.stringify(user),
        })
            .then((res) => {
                if (res.status !== 200) throw new Error(res.status);
                else return res.json();
            })
            .then((result) => {
                localStorage.setItem("token", result.access_token);
                history.push("/");
            })
            .catch(err => {
                setFailed(true);
            })
    };

    return (
        <Fragment>
            <Container>
                <Row>
                    <Col lg={6} md={6} sm={12}>
                        <div className="login-box p-5">
                            <h3><FontAwesomeIcon size="2x" icon={faUserCircle}/></h3>
                            <Alert variant="danger" show={failed} className="login-alert">Tên đăng nhập hoặc mật khẩu không chính xác</Alert>
                            <Form>
                                <Form.Group className="mb-3">
                                    <Form.Control
                                        type="input"
                                        placeholder="Tên đăng nhập"
                                        onChange={(event) => {
                                            setUser({ ...user, username: event.target.value });
                                        }}
                                    />
                                </Form.Group>

                                <Form.Group className="mb-3">
                                    <Form.Control
                                        type="password"
                                        placeholder="Mật Khẩu"
                                        onChange={(event) => {
                                            setUser({ ...user, password: event.target.value });
                                        }}
                                    />
                                </Form.Group>
                                <Button
                                    variant="primary"
                                    className="btn-submit"
                                    onClick={onSubmitHandler}
                                >
                                    Đăng nhập
                                </Button>
                            </Form>
                        </div>
                    </Col>
                </Row>
            </Container>
        </Fragment>
    );
};

export default LoginForm;
