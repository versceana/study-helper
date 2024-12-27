import { useState, useEffect } from 'react';

const TaskList = () => {
  const savedTasks = JSON.parse(localStorage.getItem('tasks')) || [];
  const savedIsDarkMode = localStorage.getItem('isDarkMode') === 'true';

  const [tasks, setTasks] = useState(savedTasks);
  const [newTask, setNewTask] = useState('');
  const [editingTaskId, setEditingTaskId] = useState(null);
  const [editedTaskText, setEditedTaskText] = useState('');
  const [isDarkMode, setIsDarkMode] = useState(savedIsDarkMode);

  useEffect(() => {
    localStorage.setItem('tasks', JSON.stringify(tasks));
  }, [tasks]);

  useEffect(() => {
    localStorage.setItem('isDarkMode', isDarkMode);
  }, [isDarkMode]);

  const toggleTheme = () => {
    setIsDarkMode((prev) => !prev);
  };

  const addTask = () => {
    if (newTask.trim()) {
      setTasks([...tasks, { id: Date.now(), text: newTask }]);
      setNewTask('');
    }
  };

  const deleteTask = (id) => {
    setTasks(tasks.filter(task => task.id !== id));
  };

  const startEditingTask = (id, text) => {
    setEditingTaskId(id);
    setEditedTaskText(text);
  };

  const saveEditedTask = (id) => {
    setTasks(tasks.map(task =>
      task.id === id ? { ...task, text: editedTaskText } : task
    ));
    setEditingTaskId(null);
    setEditedTaskText('');
  };

  return (
    <div className={`min-h-screen p-4 ${isDarkMode ? 'bg-gray-900 text-white' : 'bg-gray-100 text-black'}`}>
      <h2 className="text-2xl font-bold mb-4">Список задач</h2>
      <div className="flex mb-4">
        <input
          type="text"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          className={`flex-1 p-2 border rounded ${
            isDarkMode ? 'bg-gray-700 border-gray-600 text-white' : 'bg-white border-gray-300 text-black'
          }`}
          placeholder="Новая задача"
        />
        <button
          onClick={addTask}
          className={`ml-2 px-4 py-2 ${
            isDarkMode ? 'bg-green-600 hover:bg-green-700' : 'bg-green-500 hover:bg-green-600'
          } text-white rounded transition duration-300`}
        >
          Добавить
        </button>
      </div>
      <ul>
        {tasks.map(task => (
          <li
            key={task.id}
            className={`flex justify-between items-center p-2 border-b ${
              isDarkMode ? 'border-gray-700' : 'border-gray-300'
            }`}
          >
            {editingTaskId === task.id ? (
              <input
                type="text"
                value={editedTaskText}
                onChange={(e) => setEditedTaskText(e.target.value)}
                className={`flex-1 p-1 border rounded ${
                  isDarkMode ? 'bg-gray-700 border-gray-600 text-white' : 'bg-white border-gray-300 text-black'
                }`}
              />
            ) : (
              <span>{task.text}</span>
            )}
            <div className="space-x-2">
              {editingTaskId === task.id ? (
                <button
                  onClick={() => saveEditedTask(task.id)}
                  className={`px-2 py-1 ${
                    isDarkMode ? 'bg-blue-600 hover:bg-blue-700' : 'bg-blue-500 hover:bg-blue-600'
                  } text-white rounded transition duration-300`}
                >
                  Сохранить
                </button>
              ) : (
                <button
                  onClick={() => startEditingTask(task.id, task.text)}
                  className={`px-2 py-1 ${
                    isDarkMode ? 'bg-yellow-600 hover:bg-yellow-700' : 'bg-yellow-500 hover:bg-yellow-600'
                  } text-white rounded transition duration-300`}
                >
                  Редактировать
                </button>
              )}
              <button
                onClick={() => deleteTask(task.id)}
                className={`px-2 py-1 ${
                    isDarkMode ? 'bg-red-600 hover:bg-red-700' : 'bg-red-500 hover:bg-red-600'
                  } text-white rounded transition duration-300`}
              >
                Удалить
              </button>
            </div>
          </li>
        ))}
      </ul>
      <button
        onClick={toggleTheme}
        className={`fixed bottom-4 right-4 px-6 py-2 ${
          isDarkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-200 hover:bg-gray-300'
        } text-white rounded-full transition duration-300`}
      >
        {isDarkMode ? '🌙' : '☀️'}
      </button>
    </div>
  );
};

export default TaskList;