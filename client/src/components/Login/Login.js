import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Login.css';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  // Handle form submission
  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:3001/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Login failed!');
      }

      const data = await response.json();
      const token = data.token;

      // Store the token securely (localStorage for simplicity)
      localStorage.setItem('jwt', token);

      // Redirect to dashboard or home page
      navigate('/user');
    } catch (error) {
      alert(`Login error: ${error.message}`);
    }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleLogin} className="form">
        <div className="inputGroup">
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="input"
            required
          />
        </div>
        <button type="submit" className="button">Login</button>
      </form>
      <p>
        Don't have an account? <a href="/register">Register here</a>
      </p>
    </div>
  );
};

export default Login;
