import styles from "./Login.module.scss";

function Login(props) {
  return (
    <div className={styles.Login}>
      <form className={styles.Login_Form}>
        <h2>Vector</h2>
        <input
          type="text"
          name="Username"
          placeholder="Username"
          id="Username"
        />
        <input type="password" name="Password" placeholder="Password" id="" />
        <div>
          <input type="checkbox" name="Remember" placeholder="Remember" id="" />
          <label htmlFor="Remeber">Remember me</label>
        </div>
        <input type="submit" value="Login" />
      </form>
    </div>
  );
}

export default Login;
