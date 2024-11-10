import styles from "./styles.module.scss"
const Header = () => {
    return (
        <div className={styles.header}>
            <svg xmlns="http://www.w3.org/2000/svg" width="52" height="52" viewBox="0 0 52 52" fill="none">
                <circle cx="26" cy="26" r="26" fill="#D9D9D9" />
            </svg>
            <p>Пользователь</p>
        </div>
    );
}

export { Header };