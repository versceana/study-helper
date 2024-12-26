import React, { useState, useEffect } from 'react';
import { GoogleLogin, GoogleLogout } from 'react-google-login';
import { listEvents, addEvent } from '../services/api';

const Calendar = () => {
  const [events, setEvents] = useState([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const clientId = process.env.REACT_APP_GOOGLE_CLIENT_ID;

  const onLoginSuccess = (response) => {
    console.log('Успешный вход:', response);
    setIsLoggedIn(true);
    fetchEvents(response.tokenId);
  };

  const onLoginFailure = (error) => {
    console.error('Ошибка входа:', error);
  };

  const onLogoutSuccess = () => {
    console.log('Успешный выход');
    setIsLoggedIn(false);
    setEvents([]);
  };

  const fetchEvents = async (token) => {
    try {
      const events = await listEvents(token);
      setEvents(events);
    } catch (error) {
      console.error('Ошибка при загрузке событий:', error);
    }
  };

  const handleAddEvent = async (event) => {
    try {
      const token = localStorage.getItem('google_token');
      await addEvent(token, event);
      fetchEvents(token);
    } catch (error) {
      console.error('Ошибка при добавлении события:', error);
    }
  };

  return (
    <div className="p-4 bg-white dark:bg-gray-800 rounded-lg shadow">
      <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-4">
        Календарь
      </h3>
      {!isLoggedIn ? (
        <GoogleLogin
          clientId={clientId}
          buttonText="Войти через Google"
          onSuccess={onLoginSuccess}
          onFailure={onLoginFailure}
          cookiePolicy={'single_host_origin'}
        />
      ) : (
        <GoogleLogout
          clientId={clientId}
          buttonText="Выйти"
          onLogoutSuccess={onLogoutSuccess}
        />
      )}
      <div className="mt-4">
        {events.map((event) => (
          <div key={event.id} className="p-2 border-b">
            <p className="text-gray-900 dark:text-white">{event.summary}</p>
            <p className="text-gray-600 dark:text-gray-400">
              {event.start.dateTime} - {event.end.dateTime}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Calendar;