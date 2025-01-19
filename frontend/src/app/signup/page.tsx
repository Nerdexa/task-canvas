'use client';
import { memo, forwardRef, ReactNode, useState } from 'react';

import { useRouter } from 'next/navigation';

import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import LockIcon from '@mui/icons-material/Lock';
import MailOutlineIcon from '@mui/icons-material/MailOutline';
import CircularProgress from '@mui/material/CircularProgress';

import { useForm, SubmitHandler, Controller, SubmitErrorHandler } from 'react-hook-form';

import { useSnackbar } from '@/_components/contexts/SnackbarContext';
import Box from '@/_components/mui/Box';
import Button from '@/_components/mui/Button';
import Stack from '@/_components/mui/Stack';
import RegistrationFormBox from '@/_components/organisms/RegistrationFormBox';
import useSignUp from '@/hooks/useSignUp';
import TextFieldWithIcon from '@/_components/molecules/TextFieldWithIcon';
type InputProps = {
  email: string;
  password: string;
};

const SignUp = () => {
  const router = useRouter();
  const { showError, showSuccess } = useSnackbar();
  const { signUp } = useSignUp();
  const { control, handleSubmit } = useForm<InputProps>({
    defaultValues: {
      email: '',
      password: '',
    },
  });
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit: SubmitHandler<InputProps> = async (values) => {
    try {
      setIsLoading(true);
      await signUp(values);
      router.push('/signin');
      showSuccess('アカウント作成に成功しました');
    } catch (e) {
      showError('アカウント作成に失敗しました');
      console.error(e);
    } finally {
      setIsLoading(false);
    }
  };

  const onSubmitError: SubmitErrorHandler<InputProps> = (errors) => {};

  return (
    <RegistrationFormBox description="ユーザー情報を入力してください。">
      <Stack
        component="form"
        onSubmit={handleSubmit(onSubmit)}
      >
        <Controller
          name="email"
          control={control}
          rules={{ required: 'メールアドレスを入力してください' }}
          render={({ field }) => (
            <TextFieldWithIcon
              {...field}
              label="メール"
              placeholder="メールを入力してください"
              type="email"
              required
              disabled={isLoading}
              icon={
                <MailOutlineIcon
                  sx={{ color: 'icon.blue', height: 20, wight: 20, marginRight: 1 }}
                />
              }
            />
          )}
        />
        <Controller
          name="password"
          control={control}
          render={({ field }) => (
            <TextFieldWithIcon
              {...field}
              label="パスワード"
              type="password"
              placeholder="パスワードを入力してください"
              disabled={isLoading}
              icon={<LockIcon sx={{ color: 'icon.blue', height: 20, wight: 20, marginRight: 1 }} />}
            />
          )}
        />
        {isLoading ? (
          <Box
            component={'div'}
            aria-busy={isLoading}
            sx={{ width: '100%', display: 'flex', justifyContent: 'center' }}
          >
            <CircularProgress size={32} />
          </Box>
        ) : (
          <Button
            variant="contained"
            sx={{ marginTop: 2 }}
            disabled={isLoading}
            type="submit"
          >
            アカウントを作成する
          </Button>
        )}
      </Stack>
      <Box sx={{ mt: 1.5, display: 'flex', justifyContent: 'flex-start' }}>
        <Button
          sx={{ width: 'fit-content', fontSize: 12 }}
          variant="text"
          onClick={() => {
            router.push('/signin');
          }}
        >
          <ArrowBackIcon sx={{ color: 'icon.blue', height: 16, wight: 16, mr: 1 }} />
          サインイン
        </Button>
      </Box>
    </RegistrationFormBox>
  );
};

export default SignUp;
