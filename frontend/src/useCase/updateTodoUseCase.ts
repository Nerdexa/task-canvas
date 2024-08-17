import { Todo } from '@/domain/todo';
import TodoGateway from '@/gateway/todo';

export const updateTodo = async (id: Todo['id'], completed: Todo['completed']): Promise<void> => {
  const todoGateway = new TodoGateway();

  await todoGateway.updateTodoGateway(id, completed);
};
