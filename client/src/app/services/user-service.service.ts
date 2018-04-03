import { Injectable } from '@angular/core';
import { HttpClientModule, HttpClient } from '@angular/common/http';

@Injectable()
export class UserServiceService {

  constructor(private http: HttpClient) { }

  getUsers(test = false) {
    if (test === false) {
      return this.http.get('http://localhost:8000/user');
    } else {
      return this.http.get('http://localhost:8000/firstUser');
    }
  }
}
