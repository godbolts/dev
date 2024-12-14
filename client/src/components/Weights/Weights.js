import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom"; // Import useNavigate
import axios from "axios";
import './Weights.css';

const PreferenceWeights = () => {

  const token = localStorage.getItem('jwt');
  // State to hold weights
  const [weights, setWeights] = useState({
    distance: 0,
    age: 0,
    food: 0,
    hobbies: 0,
    music: 0,
  });

  const [error, setError] = useState("");
  const [loading, setLoading] = useState({}); // State for button loading
  const [isFetching, setIsFetching] = useState(false); // State for fetch loading

  const navigate = useNavigate(); // Initialize useNavigate for navigation

  // Fetch weights on component mount
  useEffect(() => {
    fetchWeights();
  }, []);

  const fetchWeights = async () => {
    setIsFetching(true); // Set fetching state to true
    try {
      const response = await axios.get("http://localhost:3001/api/wigh/get", {
        headers: {
          Authorization: `Bearer ${token}`, 
          'Content-Type': 'application/json',
        }
      });

      // Transform backend keys (snake_case) to frontend keys (camelCase)
      const transformedWeights = {
        distance: response.data.weigh_distance,
        age: response.data.weigh_age,
        food: response.data.weigh_food,
        hobbies: response.data.weigh_hobbies,
        music: response.data.weigh_music,
      };

      setWeights(transformedWeights);
      setError(""); // Clear any existing error
    } catch (err) {
      setError("Failed to fetch weights. Please try again.");
      console.error(err);
    }
    setIsFetching(false); // Set fetching state to false
  };

  const updateWeight = async (type, value) => {
    const endpoints = {
      distance: "http://localhost:3001/api/wigh/dist",
      age: "http://localhost:3001/api/wigh/age",
      food: "http://localhost:3001/api/wigh/food",
      hobbies: "http://localhost:3001/api/wigh/hobby",
      music: "http://localhost:3001/api/wigh/music",
    };

    const url = endpoints[type];

    if (!url) {
      setError(`Invalid type: ${type}`);
      console.error(`Invalid type: ${type}`);
      return;
    }

    setLoading((prev) => ({ ...prev, [type]: true })); // Set button loading state
    try {
      const payload = { number: value }; // Payload for the backend
      await axios.post(url, payload, {
        headers: {
          Authorization: `Bearer ${token}`, 
          'Content-Type': 'application/json', 
        }
      });
      setWeights({ ...weights, [type]: value }); // Update state
      setError(""); // Clear any existing error
    } catch (err) {
      setError(`Failed to update ${type} weight. Please try again.`);
      console.error(err);
    }
    setLoading((prev) => ({ ...prev, [type]: false })); // Clear button loading state
  };

  const handleInputChange = (e, type) => {
    const value = Math.max(0, Math.min(10, parseInt(e.target.value, 10) || 0)); // Clamp between 0 and 10
    setWeights({ ...weights, [type]: value }); // Update local state
  };

  return (
    <div>
      <h1>Preference Weights</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {isFetching ? (
        <p>Loading weights...</p> // Display loading state
      ) : (
        <div style={{ display: "flex", flexDirection: "column", gap: "10px" }}>
          {Object.keys(weights).map((type) => (
            <div key={type} style={{ display: "flex", alignItems: "center" }}>
              <label style={{ marginRight: "10px", textTransform: "capitalize" }}>
                {type}:
              </label>
              <input
                type="number"
                min="0"
                max="10"
                value={weights[type]}
                onChange={(e) => handleInputChange(e, type)}
              />
              <button onClick={() => updateWeight(type, weights[type])} disabled={loading[type]}>
                {loading[type] ? "Saving..." : `Save ${type}`}
              </button>
            </div>
          ))}
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

export default PreferenceWeights;
