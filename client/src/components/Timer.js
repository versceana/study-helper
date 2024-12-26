import React, { useState, useEffect } from 'react';

const Timer = () => {
  const [workTime, setWorkTime] = useState(45 * 60);
  const [breakTime, setBreakTime] = useState(15 * 60);
  const [timeLeft, setTimeLeft] = useState(workTime);
  const [isActive, setIsActive] = useState(false);
  const [isWorkMode, setIsWorkMode] = useState(true);
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);
  const [isDarkMode, setIsDarkMode] = useState(false);

  useEffect(() => {
    const savedWorkTime = localStorage.getItem('workTime');
    const savedBreakTime = localStorage.getItem('breakTime');
    const savedTimeLeft = localStorage.getItem('timeLeft');
    const savedIsActive = localStorage.getItem('isActive');
    const savedIsWorkMode = localStorage.getItem('isWorkMode');
    const savedIsDarkMode = localStorage.getItem('isDarkMode');

    if (savedWorkTime) setWorkTime(parseInt(savedWorkTime));
    if (savedBreakTime) setBreakTime(parseInt(savedBreakTime));
    if (savedTimeLeft) setTimeLeft(parseInt(savedTimeLeft));
    if (savedIsActive) setIsActive(savedIsActive === 'true');
    if (savedIsWorkMode) setIsWorkMode(savedIsWorkMode === 'true');
    if (savedIsDarkMode) setIsDarkMode(savedIsDarkMode === 'true');
  }, []);

  useEffect(() => {
    localStorage.setItem('workTime', workTime);
    localStorage.setItem('breakTime', breakTime);
    localStorage.setItem('timeLeft', timeLeft);
    localStorage.setItem('isActive', isActive);
    localStorage.setItem('isWorkMode', isWorkMode);
    localStorage.setItem('isDarkMode', isDarkMode);
  }, [workTime, breakTime, timeLeft, isActive, isWorkMode, isDarkMode]);

  useEffect(() => {
    let interval = null;

    if (isActive && timeLeft > 0) {
      interval = setInterval(() => {
        setTimeLeft((prev) => prev - 1);
      }, 1000);
    } else if (isActive && timeLeft === 0) {
      clearInterval(interval);
      switchMode();
    } else {
      clearInterval(interval);
    }

    return () => clearInterval(interval);
  }, [isActive, timeLeft]);

  const switchMode = () => {
    if (isWorkMode) {
      setIsWorkMode(false);
      setTimeLeft(breakTime);
      sendNotification('☕ Время отдыха!');
    } else {
      setIsWorkMode(true);
      setTimeLeft(workTime);
      sendNotification('⏳ Время работы!');
    }
  };

  const sendNotification = (message) => {
    if (Notification.permission === 'granted') {
      new Notification(message);
    } else if (Notification.permission !== 'denied') {
      Notification.requestPermission().then((permission) => {
        if (permission === 'granted') {
          new Notification(message);
        }
      });
    }
  };

  const toggleTimer = () => {
    setIsActive(!isActive);
  };

  const resetTimer = () => {
    setIsActive(false);
    setTimeLeft(isWorkMode ? workTime : breakTime);
  };

  const handleSaveSettings = (e) => {
    e.preventDefault();
    const newWorkTime = parseInt(e.target.workTime.value) * 60;
    const newBreakTime = parseInt(e.target.breakTime.value) * 60;
    setWorkTime(newWorkTime);
    setBreakTime(newBreakTime);
    setTimeLeft(isWorkMode ? newWorkTime : newBreakTime);
    setIsSettingsOpen(false);
  };

  const toggleTheme = () => {
    setIsDarkMode(!isDarkMode);
  };

  const formatTime = (time) => {
    const minutes = Math.floor(time / 60);
    const seconds = time % 60;
    return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
  };

  return (
    <div className={`min-h-screen flex items-center justify-center ${isDarkMode ? 'bg-black text-white' : 'bg-gray-900 text-white'}`}>
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">{isWorkMode ? '⏳ Время работы' : '☕ Время отдыха'}</h1>
        <div className="text-8xl font-bold mb-8">{formatTime(timeLeft)}</div>
        <div className="space-x-4">
          <button
            onClick={toggleTimer}
            className={`px-6 py-2 ${isDarkMode ? 'bg-black hover:bg-gray-800' : 'bg-gray-800 hover:bg-gray-700'} rounded-full`}
          >
            {isActive ? 'Пауза' : 'Старт'}
          </button>
          <button
            onClick={resetTimer}
            className={`px-6 py-2 ${isDarkMode ? 'bg-black hover:bg-gray-800' : 'bg-gray-800 hover:bg-gray-700'} rounded-full`}
          >
            Сброс
          </button>
          <button
            onClick={() => setIsSettingsOpen(true)}
            className={`px-6 py-2 ${isDarkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-700 hover:bg-gray-600'} rounded-full`}
          >
            ⚙️
          </button>
        </div>
        <button
          onClick={toggleTheme}
          className="mt-8 px-6 py-2 bg-gray-800 text-white rounded-full hover:bg-gray-700"
        >
          {isDarkMode ? 'Светлая тема' : 'Темная тема'}
        </button>
      </div>

      {isSettingsOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className={`p-8 rounded-lg ${isDarkMode ? 'bg-black' : 'bg-gray-900'}`}>
            <h2 className="text-2xl font-bold mb-4">Настройки</h2>
            <form onSubmit={handleSaveSettings}>
              <div className="space-y-4">
                <label className="block">
                  Время работы (минуты):
                  <input
                    type="number"
                    name="workTime"
                    defaultValue={workTime / 60}
                    className={`w-full p-2 border ${isDarkMode ? 'bg-black border-gray-800' : 'bg-gray-900 border-gray-800'} rounded`}
                  />
                </label>
                <label className="block">
                  Время отдыха (минуты):
                  <input
                    type="number"
                    name="breakTime"
                    defaultValue={breakTime / 60}
                    className={`w-full p-2 border ${isDarkMode ? 'bg-black border-gray-800' : 'bg-gray-900 border-gray-800'} rounded`}
                  />
                </label>
              </div>
              <div className="mt-6 space-x-4">
                <button type="submit" className="px-6 py-2 bg-blue-500 text-white rounded-full hover:bg-blue-600">
                  Сохранить
                </button>
                <button
                  type="button"
                  onClick={() => setIsSettingsOpen(false)}
                  className="px-6 py-2 bg-gray-500 text-white rounded-full hover:bg-gray-600"
                >
                  Отмена
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default Timer;

