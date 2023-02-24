import { TextField, TextFieldProps } from "@mui/material";
import { ChangeEvent, ComponentProps, useState } from "react";

interface FormState {
  [field: string]: any;
}

type EditorComponent<Type, Props> = React.ComponentType<{
  value: Type;
  onChange: (value: Type | ChangeEvent) => void;
}> &
  Props;

type CustomComponent<State extends FormState> = {
  [key in keyof State]?: EditorComponent<State[key], any>;
};

type Validator<Type> = (value: Type) => string | undefined;
type Validators<State extends FormState> = {
  [key in keyof State]?: Validator<State[key]>;
};

type Labels<State> = {
  [key in keyof State]?: string;
};

type EditorProps<
  State extends FormState,
  FormComponent extends CustomComponent<State>
> = {
  [key in keyof FormComponent]: Omit<
    ComponentProps<FormComponent[key]>,
    "value" | "onChange"
  >;
};

export default function createForm<
  State extends FormState,
  FormComponents = CustomComponent<State>
>(
  initialState: State,
  config: {
    customComponents?: FormComponents;
    validators?: Validators<State>;
    labels?: Labels<State>;
  } = {}
) {
  type Props = {
    submitTitle?: string;
    onSubmit(state?: State): void;
    editorProps?: EditorProps<State, FormComponents>;
  };

  function Form(props: Props) {
    const fields = [];
    const [state, setState] = useState(initialState);

    for (const key in initialState) {
    }
  }
}

const LoginForm = createForm(
  {
    username: "",
    password: "",
  },
  {
    customComponents: {},
  }
);
