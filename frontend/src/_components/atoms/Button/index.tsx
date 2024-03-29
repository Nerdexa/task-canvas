import React from 'react';

import MuiButton, { ButtonProps as MuiButtonProps } from '../../mui/Button';

export type ButtonProps = Pick<MuiButtonProps, 'children' | 'sx' | 'onClick'>;

const Button: React.FC<ButtonProps> = ({ children, sx, onClick }) => {
  return (
    <MuiButton
      onClick={onClick}
      variant="contained"
      sx={[
        {
          textTransform: 'none',
        },
        ...(Array.isArray(sx) ? sx : [sx]),
      ]}
    >
      {children}
    </MuiButton>
  );
};

export default Button;
