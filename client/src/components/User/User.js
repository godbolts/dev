import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './User.css';

const User = () => {
  const navigate = useNavigate();
  const [userData, setUserData] = useState(null);  // State to store user data
  const [error, setError] = useState(null);  // State for any errors

  useEffect(() => {
    const token = localStorage.getItem('jwt');
    if (!token) {
      navigate('/login'); // Redirect to login if not authenticated
    }

    // Log the request details
    console.log('Making request to /api/user');
    console.log('Authorization Token:', token); // Log the token for debugging

    // Fetch user data from backend
    const fetchUserData = async () => {
      try {
        const response = await fetch('http://localhost:3001/api/user', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

                // Log the request and response details
                console.log('Request sent to /api/user');
                console.log('Request Headers:', {
                  'Authorization': `Bearer ${token}`,
                });
        if (!response.ok) {
          throw new Error('Failed to fetch user data');
        }
        const data = await response.json();
        setUserData(data);  // Set the user data
      } catch (err) {
        setError(err.message);
        console.error('Error fetching user data:', err);
      }
    };

    fetchUserData();
    
    // Store browser location in a global variable when the component mounts
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;
          // Save the location in a global variable (e.g., window object)
          window.userLocation = { latitude, longitude };
          console.log('Browser location saved:', window.userLocation);
        },
        (err) => {
          console.error('Error getting location:', err);
        }
      );
    } else {
      console.log('Geolocation not supported');
    }
  }, [navigate]);

  const handleNavigation = (path) => {
    navigate(path);
  };

  const handleEditProfile = () => {
    navigate('/profileedit');
  };

  const handleEditBio = () => {
    navigate('/bioedit');
  };

  const handleEditPreferences = () => {
    navigate('/preferenceedit');
  };

  const handleEditWeights = () => {
    navigate('/weightedit');
  };

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      {/* Navigation Buttons */}
      <div>
        <button onClick={() => handleNavigation('/recommendations')}>Recommendations</button>
        <button onClick={() => handleNavigation('/connections')}>Connections</button>
        <button onClick={() => handleNavigation('/messages')}>Messages</button>
      </div>

      {/* User Info */}
      {userData && (
        <div>
          <h2>{userData.username}</h2>
          <div>
            <img
              src={userData.profilePicture}
              alt="Profile"
              style={{ width: '150px', height: '150px', borderRadius: '50%' }}
            />
          </div>
          <h2>{userData.firstName} {userData.middleName} {userData.lastName}</h2>
          <p>Home City: {userData.city}</p>
        </div>
      )}

      {/* Action Buttons */}
      <div>
        <button onClick={handleEditProfile}>Edit Profile</button>
        <button onClick={handleEditBio}>Edit Bio</button>
        <button onClick={handleEditPreferences}>Edit Preferences</button>
        <button onClick={handleEditWeights}>Edit Weights</button>
      </div>
    </div>
  );
};

export default User;
