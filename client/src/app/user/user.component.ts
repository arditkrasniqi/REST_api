import {Component, OnInit} from '@angular/core';
import { UserServiceService } from '../services/user-service.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})

export class UserComponent implements OnInit {
  users: any;

  constructor(private userService: UserServiceService) {
  }

  ngOnInit() {
    this.users = this.userService.getUsers(true);
    setTimeout(() => {
      this.users = this.userService.getUsers();
    }, 5000);
  }

}
