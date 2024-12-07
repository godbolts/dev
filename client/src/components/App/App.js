import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from '../Login/Login';
import RegisterPage from '../Registration/Registration';
import User from '../User/User';
import Bio from '../Bio/Bio';
import Preferences from '../Preferences/Preferences';
import Profile from '../Profile/Profile';
import Weights from '../Weights/Weights';

const App = () => {
  // Check if the user is authenticated by looking for the JWT token in localStorage
  const isAuthenticated = () => {
    const token = localStorage.getItem('jwt');
    return token ? true : false;
  };

  // Protected route component: If not authenticated, redirect to login
  const ProtectedRoute = ({ element }) => {
    return isAuthenticated() ? element : <Navigate to="/" replace />;
  };

  return (
    <Router>
      <div>
        {/* Add a layout or header component here if needed */}
        <Routes>
          {/* Login route */}
          <Route path="/" element={<LoginPage />} />
          
          {/* Register route */}
          <Route path="/register" element={<RegisterPage />} />
          
          {/* Protected dashboard route */}
          <Route path="/user" element={<ProtectedRoute element={<User />} />} />
          <Route path="/bioedit" element={<ProtectedRoute element={<Bio />} />} />
          <Route path="/preferenceedit" element={<ProtectedRoute element={<Preferences />} />} />
          <Route path="/profileedit" element={<ProtectedRoute element={<Profile />} />} />
          <Route path="/weightedit" element={<ProtectedRoute element={<Weights />} />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
