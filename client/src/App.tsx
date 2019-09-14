import React, {KeyboardEvent, useEffect, useState} from 'react';
import {
    Box,
    Button,
    Container,
    CssBaseline,
    Link,
    List,
    ListItem,
    ListItemSecondaryAction,
    ListItemText,
    makeStyles,
    TextField,
    Typography
} from "@material-ui/core";
import {todoServiceClient} from "./proto/TodoServiceClientPb";
import {
    addTodoParams,
    completeResponse,
    completeTodoParams,
    deleteResponse,
    deleteTodoParams,
    getTodoParams,
    todoObject,
    todoResponse
} from "./proto/todo_pb";
import {Error} from "grpc-web";

interface Todo {
    id: string;
    task: string;
    completed: boolean;
}

interface TodoProps {
    todo: Todo;
    classes: any;
    completeTodo: (id: string) => void;
    deleteTodo: (id: string) => void;
}

const useStyles = makeStyles(theme => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        minHeight: '100vh',
    },
    main: {
        marginTop: theme.spacing(8),
        marginBottom: theme.spacing(2),
    },
    footer: {
        padding: theme.spacing(2),
        marginTop: 'auto',
        backgroundColor: 'white',
    },
    button: {
        margin: theme.spacing(1),
        justifyContent: 'center',
        alignItems: 'center',
    },
    textField: {
        marginTop: 0,
        marginLeft: theme.spacing(1),
        marginRight: theme.spacing(1),
        width: 300,
        justifyContent: 'center',
    },
    list: {
        width: '100%',
        maxWidth: 660,
        backgroundColor: theme.palette.background.paper,
    },
    lineThrough: {
        textDecoration: 'line-through'
    }
}));

const Copyright = () => {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="https://barbarov.com/">
                barbarov.com
            </Link>{' '}
            {new Date().getFullYear()}
            {'.'}
        </Typography>
    );
};

const TaskItem = (props: TodoProps) => {
    return (<ListItem alignItems="flex-start">
        <ListItemText
            primary={
                <React.Fragment>
                    <Typography
                        color="textPrimary"
                        className={props.todo.completed ? props.classes.lineThrough : ''}
                    >{props.todo.task}</Typography>
                </React.Fragment>
            }
            secondary={
                <React.Fragment>
                    <Typography
                        component="span"
                        variant="body2"
                        color="textPrimary"
                        className={props.todo.completed ? props.classes.lineThrough : ''}
                    >
                        {props.todo.id}
                    </Typography>
                </React.Fragment>
            }
        />
        <ListItemSecondaryAction>
            <Button variant="contained" color="primary" className={props.classes.button}
                    onClick={() => props.completeTodo(props.todo.id)}>
                Complete
            </Button>
            <Button variant="contained" color="secondary" className={props.classes.button}
                    onClick={() => props.deleteTodo(props.todo.id)}>
                Delete
            </Button>
        </ListItemSecondaryAction>
    </ListItem>);
};

const App: React.FC = () => {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [error, setError] = useState(true);
    const [task, setTask] = useState("");
    const classes = useStyles();
    const client = new todoServiceClient("http://localhost:9090", null, null);

    useEffect(() => {
        getTodos();
    }, []);

    const keyPressed = (event: KeyboardEvent<HTMLDivElement>) => {
        if (event.key === "Enter") {
            addTodo();
        }
    };

    const getTodos = () => {
        let getRequest = new getTodoParams();
        client.getTodos(getRequest, {}, (err: Error, response: todoResponse) => {
            if (err) {
                console.error(err);
                return;
            }
            let todos = response.toObject().todosList.map(todo => ({
                id: todo.id,
                task: todo.task,
                completed: todo.completed
            }));
            setTodos(todos);
        });
    };

    const addTodo = () => {
        let request = new addTodoParams();
        request.setTask(task);
        client.addTodo(request, {}, (err: Error, response: todoObject) => {
            setTask('');
            setError(true);
            getTodos();
        });
    };

    const completeTodo = (id: string) => {
        let request = new completeTodoParams();
        request.setId(id);
        client.completeTodo(request, {}, (err: Error, response: completeResponse) => {
            if (response.getMessage() === "success") {
                getTodos();
            }
        })
    };

    const deleteTodo = (id: string) => {
        let deleteRequest = new deleteTodoParams();
        deleteRequest.setId(id);
        client.deleteTodo(deleteRequest, {}, (err: Error, response: deleteResponse) => {
            if (response.getMessage() === "success") {
                getTodos();
            }
        });
    };

    return (
        <div className={classes.root}>
            <CssBaseline/>
            <Container component="main" className={classes.main} maxWidth="lg">
                <Typography variant="h2" component="h1" gutterBottom>
                    Todo list
                </Typography>

                <Typography variant="h5" align="center">Add new tasks below.</Typography>

                <Box display="flex" alignItems="center" justifyContent="center">
                    <TextField id="standard-name"
                               label="Name"
                               margin="normal"
                               className={classes.textField}
                               value={task}
                               placeholder={task}
                               onKeyPress={(e) => keyPressed(e)}
                               onChange={e => {
                                   setTask(e.target.value);
                                   setError(e.target.value.length === 0);
                               }}
                    />
                    <Button variant="contained" color="primary" className={classes.button} onClick={() => addTodo()} disabled={error}>
                        Add task
                    </Button>
                </Box>
                <Box display="flex" flexDirection="column" alignItems="center">
                    <List className={classes.list}>
                        {todos.map(todo =>
                            <TaskItem key={todo.id} todo={todo}
                                      classes={classes}
                                      completeTodo={completeTodo}
                                      deleteTodo={deleteTodo}/>
                        )}
                    </List>
                </Box>
            </Container>
            <footer className={classes.footer}>
                <Container maxWidth="lg">
                    <Copyright/>
                </Container>
            </footer>
        </div>
    );
};

export default App;
