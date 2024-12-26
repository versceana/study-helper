import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL;

// Получение событий из Google Calendar
export const listEvents = async (token) => {
  const response = await axios.get(`${API_URL}/calendar/events`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data;
};

// Добавление события в Google Calendar
export const addEvent = async (token, event) => {
  const response = await axios.post(`${API_URL}/calendar/events`, event, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data;
};

// Пример функции для работы с задачами
export const fetchTasks = async () => {
  const response = await axios.get(`${API_URL}/tasks`);
  return response.data;
};

export const createTask = async (task) => {
  const response = await axios.post(`${API_URL}/tasks`, task);
  return response.data;
};