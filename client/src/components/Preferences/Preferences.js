import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Preferences.css';

const Preferences = () => {
  const token = localStorage.getItem('jwt'); // Get the token from localStorage
  const navigate = useNavigate(); // Initialize useNavigate for navigation
  const [preferences, setPreferences] = useState({
    food: [],
    hobby: [],
    music: [],
  });

  const [selectedPreferences, setSelectedPreferences] = useState({
    food: [],
    hobby: [],
    music: [],
  });

  useEffect(() => {
    // Fetch the preferences from the backend with authorization header
    fetch('http://localhost:3001/pref/mapget', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setPreferences({
          food: data.food || [],
          hobby: data.hobby || [],
          music: data.music || [],
        });
      })
      .catch((error) => console.error('Error fetching preferences:', error));

    // Fetch the user's selected preferences from the backend with authorization header
    fetch('http://localhost:3001/pref/get', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token in the authorization header
        'Content-Type': 'application/json',
      },
    })
      .then((response) => response.json())
      .then((data) => {
        const userSelectedPreferences = {
          food: data.food || [],
          hobby: data.hobby || [],
          music: data.music || [],
        };

        setSelectedPreferences(userSelectedPreferences);
      })
      .catch((error) => console.error('Error fetching user preferences:', error));
  }, [token]);

  const handleCheckboxChange = (category, code) => {
    setSelectedPreferences((prevSelected) => {
      const newSelection = { ...prevSelected };

      if (newSelection[category].includes(code)) {
        // Uncheck the box: remove code from the list and set `true`
        newSelection[category] = newSelection[category].filter(
          (item) => item !== code
        );
        // Make the POST request to disable this preference
        postPreference(category, code, true);
      } else {
        // Check the box: add code to the list and set `false`
        newSelection[category].push(code);
        // Make the POST request to enable this preference
        postPreference(category, code, false);
      }

      return newSelection;
    });
  };

  const postPreference = (category, code, isUnchecked) => {
    let url = '';
    let data = {};

    // Define the endpoint based on the category
    if (category === 'food') {
      url = 'http://localhost:3001/pref/food';
    } else if (category === 'hobby') {
      url = 'http://localhost:3001/pref/hobby';
    } else if (category === 'music') {
      url = 'http://localhost:3001/pref/music';
    }

    // Set the data to be sent in the POST request
    data = {
      code,
      isUnchecked, // true for unchecked (disable), false for checked (enable)
    };

    // Send a POST request with the updated data and authorization header
    fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token in the authorization header
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (!response.ok) {
          console.error('Failed to update preference');
        }
      })
      .catch((error) => console.error('Error sending preference update:', error));
  };

  const renderCheckboxes = (category, options) => {
    return options.map((option) => (
      <div key={option.code} className="checkbox-item">
        <input
          type="checkbox"
          id={option.code}
          checked={selectedPreferences[category].includes(option.code)}
          onChange={() => handleCheckboxChange(category, option.code)}
        />
        <label htmlFor={option.code}>{option.description}</label>
      </div>
    ));
  };

  return (
    <div className="preferences-container">
      <h1>User Preferences</h1>
      <div className="preferences-category">
        <h2>Food</h2>
        {renderCheckboxes('food', preferences.food)}
      </div>

      <div className="preferences-category">
        <h2>Hobbies</h2>
        {renderCheckboxes('hobby', preferences.hobby)}
      </div>

      <div className="preferences-category">
        <h2>Music</h2>
        {renderCheckboxes('music', preferences.music)}
      </div>

      {/* Optionally, you can display the selected preferences */}
      <div className="selected-preferences">
        <h3>Selected Preferences</h3>
        <pre>{JSON.stringify(selectedPreferences, null, 2)}</pre>
      </div>
      <button
        onClick={() => navigate("/user")} // Navigate to the user page
      >
        Return to User Page
      </button>
    </div>
  );
};

export default Preferences;
