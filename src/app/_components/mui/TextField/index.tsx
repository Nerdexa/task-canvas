import React from 'react';

import { TextField as MuiTextField, TextFieldProps as MuiTextFieldProps } from '@mui/material';

export type TextFieldProps = MuiTextFieldProps;

const TextField: React.FC<TextFieldProps> = (props: TextFieldProps) => {
  return <MuiTextField {...props} />;
};

export default TextField;
