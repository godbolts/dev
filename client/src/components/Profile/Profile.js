import React, { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';
import axios from "axios";
import { CitySelect, CountrySelect, StateSelect } from 'react-country-state-city';
import 'react-country-state-city/dist/react-country-state-city.css';
import './Profile.css';

const ProfileSettings = () => {
  const token = localStorage.getItem('jwt');
    const [formData, setFormData] = useState({
      username: "",
      email: "",
      firstName: "",
      middleName: "",
      lastName: "",
      password: "",
      confirmPassword: "",
      city: "",
    });
  
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");
    const [loading, setLoading] = useState({});
    const [isFetching, setIsFetching] = useState(false);
  
    const [countryId, setCountryId] = useState(null);
    const [stateId, setStateId] = useState(null);
    const navigate = useNavigate(); // Initialize useNavigate for navigation
  
    const endpoints = {
      username: "http://localhost:3001/edit/user",
      email: "http://localhost:3001/edit/email",
      firstName: "http://localhost:3001/edit/first",
      middleName: "http://localhost:3001/edit/middle",
      lastName: "http://localhost:3001/edit/last",
      password: "http://localhost:3001/edit/pass",
    };
  
    useEffect(() => {
        const fetchProfile = async () => {
          setIsFetching(true);
          try {
            const response = await axios.get("http://localhost:3001/api/user", {
              headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json',
              }
            }); // Endpoint to get user details
            setFormData((prevData) => ({ ...prevData, ...response.data }));
            setError("");
          } catch (err) {
            setError("Failed to fetch profile data. Please try again.");
            console.error(err);
          }
          setIsFetching(false);
        };
      
        fetchProfile();
      }, []);
  
    const handleChange = (e) => {
      const { name, value } = e.target;
      setFormData({ ...formData, [name]: value });
    };
  
    const handleCitySelect = (city) => {
        console.log("Selected city:", city);
        if (city && city.name && city.latitude && city.longitude) {
          setFormData((prevData) => ({
            ...prevData,
            city: city.name,
            latitude: city.latitude,
            longitude: city.longitude,
          }));
        } else {
          console.error("City data is incomplete.");
        }
      };
      
      const updateField = async (field) => {
        if (field === "password" && formData.password !== formData.confirmPassword) {
          setError("Passwords do not match.");
          return;
        }
      
        setLoading((prev) => ({ ...prev, [field]: true }));
        setError("");
        setSuccess("");
      
        try {
          if (field === "city") {
            // Send city details (name, latitude, longitude) in the payload
            await axios.post("http://localhost:3001/edit/city", {
              name: formData.city,
              latitude: formData.latitude,
              longitude: formData.longitude,
            });
          } else {
            // Send single field value for other fields
            await axios.post(endpoints[field], { value: formData[field] },
              {
                headers: {
                  Authorization: `Bearer ${token}`,
                  'Content-Type': 'application/json',
                }
              }
            );
          }
          setSuccess(`${field} updated successfully!`);
        } catch (err) {
          setError(`Failed to update ${field}. Please try again.`);
          console.error(err);
        }
        setLoading((prev) => ({ ...prev, [field]: false }));
      };
  
    return (
      <div className="container">
        <h2>Edit Profile</h2>
        {error && <p className="error">{error}</p>}
        {success && <p className="success">{success}</p>}
        {isFetching ? (
          <p>Loading profile...</p>
        ) : (
          <form className="form">
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
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("username")}
                disabled={loading.username}
              >
                {loading.username ? "Saving..." : "Save"}
              </button>
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
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("email")}
                disabled={loading.email}
              >
                {loading.email ? "Saving..." : "Save"}
              </button>
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
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("firstName")}
                disabled={loading.firstName}
              >
                {loading.firstName ? "Saving..." : "Save"}
              </button>
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
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("middleName")}
                disabled={loading.middleName}
              >
                {loading.middleName ? "Saving..." : "Save"}
              </button>
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
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("lastName")}
                disabled={loading.lastName}
              >
                {loading.lastName ? "Saving..." : "Save"}
              </button>
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
              <label>Confirm Password:</label>
              <input
                type="password"
                name="confirmPassword"
                value={formData.confirmPassword}
                onChange={handleChange}
                className="input"
                required
              />
              <button
                type="button"
                className="submitButton"
                onClick={() => updateField("password")}
                disabled={loading.password}
              >
                {loading.password ? "Saving..." : "Save"}
              </button>
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
          </form>
        )}
              <button
        onClick={() => navigate("/user")} // Navigate to the user page
      >
        Return to User Page
      </button>
      </div>
    );
  };
  
  export default ProfileSettings;


  
  
  
  
  
  
  
  
  
  









