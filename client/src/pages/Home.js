import React from 'react';
import Timer from '../components/Timer';
import TaskList from '../components/TaskList';
import Calendar from './Calendar';

const Home = () => {
  return (
    <div className="min-h-screen bg-gray-100 dark:bg-gray-900 p-4">
      <h1 className="text-3xl font-bold text-center text-gray-900 dark:text-white mb-8">
        Умный органайзер
      </h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
          <h2 className="text-2xl font-semibold text-gray-800 dark:text-white mb-4">
            Таймер
          </h2>
          <Timer />
        </div>
        <div>
          <h2 className="text-2xl font-semibold text-gray-800 dark:text-white mb-4">
            Список задач
          </h2>
          <TaskList />
        </div>
      </div>
      <div className="mt-8">
        <h2 className="text-2xl font-semibold text-gray-800 dark:text-white mb-4">
          Календарь
        </h2>
        <Calendar />
      </div>
    </div>
  );
};

export default Home;