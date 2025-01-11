import { useState, useEffect } from 'react';

import { Todo } from '@/domain/todo';
import { createTodo } from '@/useCase/createTodoUseCase';
import { getTodos } from '@/useCase/getTodoUseCase';
import { updateTodo as updateTodoUseCase } from '@/useCase/updateTodoUseCase';

export const useTodo = () => {
  const [todos, setTodos] = useState<Todo[]>([]);

  const addTodo = (content: Todo['content']) => {
    createTodo(content).then((result) => {
      setTodos([...todos, result]);
    });
  };

  const updateTodo = (id: Todo['id'], content: Todo['content'], completed: Todo['completed']) => {
    updateTodoUseCase(id, content, completed).then(() => {
      setTodos((currentTodos) => {
        return currentTodos.map((todo) => {
          if (id === todo.id) {
            const newTodo: Todo = {
              id: todo.id,
              content: content,
              completed: completed,
            };

            return newTodo;
          }

          return todo;
        });
      });
    });
  };

  useEffect(() => {
    getTodos().then((result) => {
      setTodos(result);
    });
  }, []);

  return { todos, addTodo, updateTodo };
};
