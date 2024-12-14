import React, { useEffect, useState} from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './Bio.css';

const BiographicalPage = () => {
  const token = localStorage.getItem('jwt'); // Get the token from localStorage
  // State for the biography, birth date, and age
  const [about, setAbout] = useState("");
  const [birthday, setBirthday] = useState("");
  const [age, setAge] = useState(null);

  const [newAbout, setNewAbout] = useState("");
  const [newBirthday, setNewBirthday] = useState("");
  const navigate = useNavigate(); // Initialize useNavigate for navigation
  
  // States for feedback and loading
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  // Fetch the initial biography and birthday/age on component mount
  useEffect(() => {
    const fetchBiographyData = async () => {
      setLoading(true);
      try {
        // Add the Authorization header with the Bearer token
        const aboutResponse = await axios.get("http://localhost:3001/api/biog/aboutget", {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        });

        const birthdayResponse = await axios.get("http://localhost:3001/api/biog/birthdayget", {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        });

        setAbout(aboutResponse.data.about);
        setBirthday(birthdayResponse.data.birthday);
        setAge(calculateAge(birthdayResponse.data.birthday)); // Assuming birthday is a date string

        setError("");
      } catch (err) {
        setError("Failed to fetch biographical data. Please try again.");
        console.error(err);
      }
      setLoading(false);
    };

    fetchBiographyData();
  }, [token]);

  // Helper function to calculate age from the birth date
  const calculateAge = (birthday) => {
    const birthDate = new Date(birthday);
    const ageDiff = Date.now() - birthDate.getTime();
    const ageDate = new Date(ageDiff);
    return Math.abs(ageDate.getUTCFullYear() - 1970); // Convert milliseconds to years
  };

  // Handle changes in the biography text
  const handleAboutChange = (e) => {
    setNewAbout(e.target.value);
  };

  // Handle changes in the birthday date
  const handleBirthdayChange = (e) => {
    setNewBirthday(e.target.value);
  };

  // Handle POST request to update biography
  const updateAbout = async () => {
    if (newAbout.length > 10000) {
      setError("Biography text exceeds the 10000 character limit.");
      return;
    }

    setLoading(true);
    setError("");
    setSuccess("");

    try {
      // Add the Authorization header with the Bearer token
      await axios.post(
        "http://localhost:3001/api/biog/about", 
        { newAbout: newAbout }, // Pass the newAbout inside the request body object
        {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json', // Ensure this is set for JSON data
          }
        }
      );
      setSuccess("Biography updated successfully!");
    } catch (err) {
      console.log(Request.body)
      console.log(err); // Log the actual error object for better debugging
      setError("Failed to update biography. Please try again.");
    }
    setLoading(false);
  };

  // Handle POST request to update birthday
  const updateBirthday = async () => {
    if (!newBirthday) {
      setError("Please select a valid date.");
      return;
    }

    setLoading(true);
    setError("");
    setSuccess("");

    try {
      // Add the Authorization header with the Bearer token
      await axios.post("http://localhost:3001/api/biog/birthday", { birthday: newBirthday }, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });
      setSuccess("Birthday updated successfully!");
      setBirthday(newBirthday); // Update the birthday state
      setAge(calculateAge(newBirthday)); // Recalculate age
    } catch (err) {
      setError("Failed to update birthday. Please try again.");
      console.error(err);
    }
    setLoading(false);
  };

  return (
    <div className="container">
      <h2>Biographical Information</h2>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <div>
          {/* Display current biography */}
          <div className="bio-section">
            <h3>About You</h3>
            <textarea
              value={newAbout || about}
              onChange={handleAboutChange}
              placeholder="Write about yourself"
              maxLength={10000}
            />
            <button onClick={updateAbout} disabled={loading}>
              {loading ? "Saving..." : "Save Biography"}
            </button>
          </div>

          {/* Display current birthday and age */}
          <div className="birthday-section">
            <h3>Birthday</h3>
            <p>Your current birthday: {birthday}</p>
            <p>Your current age: {age}</p>

            <input
              type="date"
              value={newBirthday || birthday}
              onChange={handleBirthdayChange}
            />
            <button onClick={updateBirthday} disabled={loading}>
              {loading ? "Saving..." : "Save Birthday"}
            </button>
          </div>

          {/* Error and success messages */}
          {error && <p className="error">{error}</p>}
          {success && <p className="success">{success}</p>}
        </div>
      )}
      <button
        onClick={() => navigate("/user")} // Navigate to the user page
      >
        Return to User Page
      </button>
    </div>
  );
};

export default BiographicalPage;
