import { useState } from 'react';

const TaskList = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');

  const addTask = () => {
    if (newTask.trim()) {
      setTasks([...tasks, { id: Date.now(), text: newTask }]);
      setNewTask('');
    }
  };

  const deleteTask = (id) => {
    setTasks(tasks.filter(task => task.id !== id));
  };

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Список задач</h2>
      <div className="flex mb-4">
        <input
          type="text"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          className="flex-1 p-2 border rounded"
          placeholder="Новая задача"
        />
        <button
          onClick={addTask}
          className="ml-2 px-4 py-2 bg-green-500 text-white rounded"
        >
          Добавить
        </button>
      </div>
      <ul>
        {tasks.map(task => (
          <li key={task.id} className="flex justify-between items-center p-2 border-b">
            <span>{task.text}</span>
            <button
              onClick={() => deleteTask(task.id)}
              className="text-red-500"
            >
              Удалить
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TaskList;