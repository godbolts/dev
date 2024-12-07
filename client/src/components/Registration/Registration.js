import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Registration.css';
import { CitySelect, CountrySelect, StateSelect } from 'react-country-state-city';
import 'react-country-state-city/dist/react-country-state-city.css';

const RegisterPage = () => {
  const [countryId, setCountryId] = useState(0);
  const [stateId, setStateId] = useState(0);
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    firstName: '',
    middleName: '',
    lastName: '',
    password: '',
    confirmPassword: '',
    city: '',
    latitude: '',
    longitude: '',
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleCitySelect = (city) => {
    console.log('Selected city:', city);
    if (city && city.name && city.latitude && city.longitude) {
      setFormData((prevData) => ({
        ...prevData,
        city: city.name,
        latitude: city.latitude,
        longitude: city.longitude,
      }));
    } else {
      console.error('City data is incomplete.');
    }
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    const {
      username,
      email,
      firstName,
      middleName,
      lastName,
      password,
      confirmPassword,
      city,
      latitude,
      longitude,
    } = formData;

    if (password !== confirmPassword) {
      alert('Passwords do not match!');
      return;
    }

    try {
      const requestBody = {
        username,
        email,
        first_name: firstName,
        middle_name: middleName,
        last_name: lastName,
        password,
        user_city: city,
        latitude: parseFloat(latitude),
        longitude: parseFloat(longitude),
      };

      const response = await fetch('http://localhost:3001/api/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Registration failed!');
      }

      const data = await response.json();

      // Assume the backend sends a token upon successful registration
      const token = data.token;

      // Store the token
      localStorage.setItem('jwt', token);

      alert(`Registration successful for ${username}.`);
      navigate('/dashboard'); // Redirect to the dashboard
    } catch (error) {
      alert(`Registration error: ${error.message}`);
    }
  };

  return (
    <div className="container">
      <h2>Register</h2>
      <form onSubmit={handleRegister} className="form">
        <div className="inputGroup">
          <label>Username:</label>
          <input
            type="text"
            name="username"
            value={formData.username}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>Email:</label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>First Name:</label>
          <input
            type="text"
            name="firstName"
            value={formData.firstName}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>Middle Name:</label>
          <input
            type="text"
            name="middleName"
            value={formData.middleName}
            onChange={handleChange}
            className="input"
          />
        </div>
        <div className="inputGroup">
          <label>Last Name:</label>
          <input
            type="text"
            name="lastName"
            value={formData.lastName}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>Password:</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <label>Confirm Password:</label>
          <input
            type="password"
            name="confirmPassword"
            value={formData.confirmPassword}
            onChange={handleChange}
            className="input"
            required
          />
        </div>
        <div className="inputGroup">
          <h6>Country</h6>
          <CountrySelect
            onChange={(e) => setCountryId(e.id)}
            placeHolder="Select Country"
          />
          <h6>State</h6>
          <StateSelect
            countryid={countryId}
            onChange={(e) => setStateId(e.id)}
            placeHolder="Select State"
          />
          <h6>City</h6>
          <CitySelect
            countryid={countryId}
            stateid={stateId}
            onChange={handleCitySelect}
            placeHolder="Select City"
          />
        </div>

        <button type="submit" className="button">Register</button>
        <p>
          Already have an account? <a href="/">Login here</a>
        </p>
      </form>
    </div>
  );
};

export default RegisterPage;
