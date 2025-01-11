import React, { useState } from 'react';

import Checkbox from '@/_components/atoms/Checkbox';
import Box from '@/_components/mui/Box';
import EditIcon from '@/_components/mui/EditIcon';
import IconButton from '@/_components/mui/IconButton';
import Typography from '@/_components/mui/Typography';

import EditPage from '../EditPage';

export type TodoCardProps = {
  id: string;
  text: string;
  checked: boolean;
  onChange: (event: React.ChangeEvent<HTMLInputElement>, checked: boolean) => void;
  onSave: (id: string, value: string) => Promise<void>;
  onDelete: (id: string) => void;
};

const TodoCard: React.FC<TodoCardProps> = ({ id, checked, onChange, text, onSave, onDelete }) => {
  const [isEditDialog, setIsEditDialog] = useState<boolean>(false);
  const [isHover, setIsHover] = useState<boolean>(false);

  const visibleEditIcon: React.MouseEventHandler<HTMLDivElement> = () => {
    setIsHover(true);
  };

  const hiddenEditIcon: React.MouseEventHandler<HTMLDivElement> = () => {
    setIsHover(false);
  };

  const handleOpenEditDialog: () => void = () => {
    setIsEditDialog(true);
  };

  const handleCloseEditDialog: () => void = () => {
    setIsEditDialog(false);
  };

  return (
    <Box
      onMouseOver={visibleEditIcon}
      onMouseLeave={hiddenEditIcon}
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        '&:hover': {
          bgcolor: 'White',
        },
      }}
    >
      <Checkbox
        checked={checked}
        onChange={onChange}
        sx={{
          '.MuiSvgIcon-root': {
            fontSize: 32,
          },
        }}
      />
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          width: '100%',
        }}
      >
        <Typography
          sx={{
            fontSize: '24px',
            marginLeft: '12px',
          }}
        >
          {text}
        </Typography>
        {isHover && (
          <>
            <IconButton
              onClick={handleOpenEditDialog}
              sx={{
                marginRight: '12px',
              }}
            >
              <EditIcon />
            </IconButton>
          </>
        )}
        <EditPage
          id={id}
          open={isEditDialog}
          onClose={handleCloseEditDialog}
          initialValue={text}
          onSave={(id, text) => {
            onSave(id, text)
              .then(() => {
                setIsEditDialog(false);
              })
              .catch(() => {
                return;
              });
          }}
          onDelete={onDelete}
        />
      </Box>
    </Box>
  );
};

export default TodoCard;
