'use client';

import { useState } from 'react';

import { useRouter } from 'next/navigation';

import Button from '@/_components/atoms/Button';
import Input, { InputProps } from '@/_components/atoms/Input';
import Title from '@/_components/molecules/Title';
import TodoCard from '@/_components/molecules/TodoCard';
import Box from '@/_components/mui/Box';
import Calender from '@/_components/mui/Calendar';
import Container from '@/_components/mui/Container';
import { Todo } from '@/domain/todo';
import useSignOut from '@/hooks/useSignOut';
import { useTodo } from '@/hooks/useTodo';

const Top = () => {
  const router = useRouter();
  const [value, setValue] = useState<string>('');
  const { todos, addTodo, updateTodo } = useTodo();
  const { signOut } = useSignOut();
  console.log(todos);
  const handleChangeText: InputProps['onChange'] = (event) => {
    const inputText = event.target.value;
    if (inputText.length > 40) return;
    setValue(inputText);
  };

  const handleCreateTodo = (value: string) => {
    addTodo(value);

    setValue('');
  };

  const handleChangeCheckbox = (
    id: Todo['id'],
    content: Todo['content'],
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    updateTodo(id, content, event.target.checked);
  };

  return (
    <Box
      sx={{
        backgroundColor: '#FFF',
        height: '100vh',
        width: '100vw',
        paddingTop: '60px',
      }}
    >
      <Button
        onClick={() => {
          signOut().then(() => {
            router.push('/signin');
          });
        }}
        sx={{
          position: 'absolute',
          top: '20px',
          right: '20px',
          color: 'white',
          '&:hover': {
            backgroundColor: 'primary.dark',
          },
        }}
      >
        ログアウト
      </Button>
      <Container maxWidth={'lg'}>
        <Box
          sx={{
            width: '100%',
            minHeight: 500,
            backgroundColor: '#F8F9FA',
            boxShadow: 7,
            borderRadius: 1,
            padding: 10,
          }}
        >
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              marginBottom: 4,
            }}
          >
            <Title href="/" />
          </Box>
          <Box
            sx={{
              position: 'relative',
              marginBottom: 3,
            }}
          >
            <Input
              value={value}
              onChange={handleChangeText}
            />
            <Calender
              sx={{
                position: 'absolute',
                top: '15px',
                right: '90px',
              }}
            />
            <Button
              sx={{
                position: 'absolute',
                top: '10px',
                right: '20px',
              }}
              disabled={!value}
              onClick={() => handleCreateTodo(value)}
            >
              Add
            </Button>
          </Box>
          <Box
            sx={{
              width: '100%',
              height: 1,
              borderBottom: 1,
              opacity: 0.1,
            }}
          />
          <Box sx={{ marginTop: 3 }}>
            {todos.map((todo, i) => {
              return (
                <TodoCard
                  id={todo.id}
                  text={todo.content}
                  checked={todo.completed}
                  onChange={(event) => {
                    handleChangeCheckbox(todo.id, todo.content, event);
                  }}
                  key={i}
                  onSave={async (id, value) => {
                    updateTodo(id, value, todo.completed);
                  }}
                  onDelete={() => {}}
                />
              );
            })}
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default Top;
