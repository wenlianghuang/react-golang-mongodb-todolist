import React, {Component, useEffect, useState} from 'react'
import axios from 'axios'
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

export default function ToDoListHook(){
    const [task,setTask] = useState('')
    const [items,SetItems] = useState([])
    let endpoint = "http://localhost:8080"
    const handleChange = (event) => {
        setTask(event.target.value)
    }
    const handleSubmit = () => {
        // let { task } = this.state;
        // console.log("pRINTING task", this.state.task);
        if (task) {
          axios
            .post(
              endpoint + "/api/task",
              {
                task,
              },
              {
                headers: {
                  "Content-Type": "application/x-www-form-urlencoded",
                },
              }
            )
            .then((res) => {
              getTask();
              setTask("")
              console.log(res);
            });
        }
    };
    useEffect(()=>{
        getTask();
    })
    const getTask = () => {
        axios.get(endpoint + "/api/task").then((res)=>{
            if(res.data){
                SetItems(res.data.map((item)=>{
                    let color = "yellow"
                    let style = {
                        wordWrap: "break-word",
                    };
                    if(item.status){
                        color = "green"
                        style["textDecorationLine"] = "line-through";
                    }
                    
                    return (
                        <Card key={item._id} color={color} fluid>
                          <Card.Content>
                            <Card.Header textAlign="left">
                              <div style={style}>{item.task}</div>
                            </Card.Header>
          
                            <Card.Meta textAlign="right">
                              <Icon
                                name="check circle"
                                color="green"
                                onClick={() => updateTask(item._id)}
                              />
                              <span style={{ paddingRight: 10 }}>Done</span>
                              <Icon
                                name="undo"
                                color="yellow"
                                onClick={() => undoTask(item._id)}
                              />
                              <span style={{ paddingRight: 10 }}>Undo</span>
                              <Icon
                                name="delete"
                                color="red"
                                onClick={() => deleteTask(item._id)}
                              />
                              <span style={{ paddingRight: 10 }}>Delete</span>
                            </Card.Meta>
                          </Card.Content>
                        </Card>
                      );
                }))
            }else{
                SetItems([]);
            }
        })
    }

    const updateTask = (id) => {
        axios
            .put(endpoint + "/api/task/" + id, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
            })
            .then((res)=>{
                console.log(res);
                getTask();
            })
    }

    const undoTask = (id) => {
        axios
          .put(endpoint + "/api/undoTask/" + id, {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          })
          .then((res) => {
            console.log(res);
            getTask();
          });
    };

    const deleteTask = (id) => {
        axios
          .delete(endpoint + "/api/deleteTask/" + id, {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          })
          .then((res) => {
            console.log(res);
            getTask();
          });
    };

    return (
        <div>
        <div className="row">
          <Header className="header" as="h2">
            TO DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={handleSubmit}>
            <Input
              type="text"
              name="task"
              onChange={handleChange}
              value={task}
              fluid
              placeholder="Create Task"
            />
            {/* <Button >Create Task</Button> */}
          </Form>
        </div>
        <div className="row">
          <Card.Group>{items}</Card.Group>
        </div>
      </div>
    )
}
