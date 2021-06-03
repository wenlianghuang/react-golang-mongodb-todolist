import React, {Component, useEffect, useState} from 'react'
import axios from 'axios'
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";
import "normalize.css"
export default function ToDoListHook(){
    const [number,setNumber] = useState('')
    const addNumber = async (e) =>{
      setNumber(e.target.value)
    }
    

    
  
    const [task,setTask] = useState('')
    const [task2,setTask2] = useState('')
    //const [task,setTask] = useState({})
    const [items,SetItems] = useState([])
    let endpoint = "https://main.d1qmmhg8f52rkl.amplifyapp.com:8090"
    const handleChange = (event) => {
        setTask(event.target.value)
    }
    const handleChange2 = (event) => {
      setTask2(event.target.value)
    }    
    const handleSubmit = async (e) => {
        // let { task } = this.state;
        // console.log("pRINTING task", this.state.task);
        e.preventDefault()
        if (task) {
          console.log("Has task")
        await axios
            .post(
              endpoint + "/api/task",
              {
                
                //userObject
                task:task,
                task2:task2,
                number:number,
              },
              {
                headers: {
                  "Content-Type": "application/x-www-form-urlencoded",
                },
              }
            )
            .then((res) => {
              console.log("Just test")
              console.log(res);
              getTask();
              setTask("")
              setTask2("")
              setNumber("")
            });
        }
    };
    useEffect(()=>{
        getTask();
    })
    const getTask = async () => {
        await axios.get(endpoint + "/api/task").then((res)=>{
            if(res.data){
                //console.log("Can show res.data: ",res.data)
                SetItems(res.data.map((item)=>{
                    let color = "yellow"
                    let style = {
                        wordWrap: "break-word",
                    };
                    if(item.status){
                        color = "green"
                        style["textDecorationLine"] = "line-through";
                    }
                    let settonumber = parseFloat(item.number)
                    return (
                        <Card key={item._id} color={color} fluid>
                          <Card.Content>
                            <Card.Header textAlign="left">
                              <div style={style}>{item.test}</div>
                              <div style={style}>{item.ttt}</div>
                              <div style={style}>{settonumber}</div>
                              {item.subtest ? 
                              <div style={style}>{item.subtest[0].substring}</div>
                              :<div style={style}></div>
                              
                            }
                            </Card.Header>
          
                            <Card.Meta textAlign="right">
                              <Icon 
                                name="cloud download"
                                color="brown"
                                onClick={()=>addSubTask(item._id)}
                              />
                              <span style={{ paddingRight: 10}}>Add</span>
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

    
    const updateTask = async (id) => {
        await axios
            .put(endpoint + "/api/doneTask/" + id, {
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
            })
            .then((res)=>{
                console.log(res);
                getTask();
            })
    }

    const undoTask = async (id) => {
        await axios
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
    const addSubTask = async (id) => {
      console.log("id: ",id);
      await axios
        .post(endpoint + "/api/task/" + id,{
          headers: {
            "Content-Type": "applicaton/x-www-form-urlencoded",
          },
        })
        .then((res) => {
          console.log(res)
          getTask();
        }) 
    }
    const deleteTask = async (id) => {
        await axios
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
            <div> 
            <Input
              type="text"
              name="task"
              onChange={handleChange}
              value={task}
              fluid
              placeholder="Create Task"
            />
            </div>
            <div>
              <Input
                type="text"
                name="task"
                onChange={handleChange2}
                value={task2}
                fluid
                placeholder=""
              />
            </div>
            <div>
              <Input
                type="text"
                name="number"
                onChange={addNumber}
                value={number}
                fluid
                placeholder=""
              />
            </div>
            
            <div>
              <input type="submit" value="Submit"/>
            </div>
            {/* <Button >Create Task</Button> */}
          </Form>
          
        </div>
        
        <div className="row">
          <Card.Group>{items}</Card.Group>
        </div>
      </div>
    )
}
