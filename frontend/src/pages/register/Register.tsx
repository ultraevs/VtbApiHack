import React, { useState } from "react";
import styles from "./styles.module.scss";
import cn from "classnames";
import { AuthLayout } from "../../layout/AuthLayout";
import { createUser } from "./http";
import { Link, useNavigate } from "react-router-dom";
import { setAuthTokenCookie } from "../../helpers"
import { RegisterInput } from "../../consts/types";
import { FaEye, FaEyeSlash } from 'react-icons/fa';
// eslint-disable-next-line no-unused-vars
React;

const Register = () => {
    const [showPassword, setShowPassword] = useState(false);
    const togglePasswordVisibility = () => {
        setShowPassword((prevShowPassword) => !prevShowPassword);
    };
    const navigate = useNavigate();
    const [loginInput, setLoginInput] = useState<RegisterInput>({
        surname: "",
        email: "",
        password: "",
    });

    const handleLoginInput = (e: any) => {
        const { name, value } = e.target;
        setLoginInput((prev: any) => ({
            ...prev,
            [name]: value,
        }));
    };

    const buttonClick = async () => {
        if (loginInput.surname !== "" && loginInput.email !== "" && loginInput.password !== "") {
            try {
                const response = await createUser(loginInput);

                if (response.success) {
                    setAuthTokenCookie(response.data);
                    window.localStorage.setItem("isAuth", "true")
                    navigate("/");
                } else {
                    alert("Ошибка: " + response.error);
                }
            } catch (error) {
                console.error(error);
            }
        } else {
            alert("Что-то не так");
        }
    };

    return (
        <AuthLayout>
            <div className={cn("container", styles.auth)}>
                <div className={styles.auth__select}>
                    <h2>Регистрация</h2>
                    <p>Регистрация банковского аккаунта</p>
                </div>
                <div className={styles.auth__form}>
                    <p>Фамилия</p>
                    <input
                        type="text"
                        name="surname"
                        value={loginInput.surname}
                        onChange={handleLoginInput}
                    />
                    <p>Эл. почта</p>
                    <input
                        type="mail"
                        name="mail"
                        value={loginInput.email}
                        onChange={handleLoginInput}
                    />
                    <p>Пароль</p>
                    <div className={styles.auth__form__password}>
                        <input
                            type={showPassword ? "text" : "password"}
                            name="password"
                            value={loginInput.password}
                            onChange={handleLoginInput}
                        />
                        <button onClick={togglePasswordVisibility}>
                            {showPassword ? <FaEye /> : <FaEyeSlash />}
                        </button>
                    </div>
                </div>
                <div className={styles.auth__actions}>
                    <div className={styles.auth__actions__memory}>
                        <input type="checkbox" name="" id="" />
                        <span>Запомнить меня</span>
                    </div>

                    <button>
                        Забыли пароль?
                    </button>
                </div>
                <div className={styles.auth__button}>
                    <button onClick={buttonClick}>Зарегистрироваться</button>
                </div>
                <div className={styles.auth__register}>
                    <span>Уже есть аккаунта? </span>
                    <Link to={"/auth"}>Войти</Link>
                </div>
            </div>
        </AuthLayout>
    );
};

export { Register };