import { useState, useEffect } from 'react';
import { Todo } from '@/domain/todo';
import { getTodos } from '@/useCase/getTodoUseCase';
import { createTodo } from '@/useCase/createTodoUseCase';
import { updateTodo as updateTodoUseCase } from '@/useCase/updateTodoUseCase';

export const useTodo = () => {
  const [todos, setTodos] = useState<Todo[]>([]);

  const addTodo = (content: Todo['content']) => {
    createTodo(content).then((result) => {
      setTodos([...todos, result]);
    });
  };

  const updateTodo = (id: Todo['id'], completed: Todo['completed']) => {
    updateTodoUseCase(id, completed).then(() => {
      const newTodos = todos.map((todo) => {
        if (id === todo.id) {
          const newTodo: Todo = {
            id: todo.id,
            content: todo.content,
            completed: completed,
          };
          return newTodo;
        }
        return todo;
      });

      setTodos(newTodos);
    });
  };

  useEffect(() => {
    getTodos().then((result) => {
      setTodos(result);
    });
  }, []);

  return { todos, addTodo, updateTodo };
};
