/*
 *
 * Copyright 2024 Alexander Sharov (@kvendingoldo).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */


/*
 *
 * NOTE: This is a sample Terraform code. Delete the context of it after the GitHub project fork.
 *
 */

variable "defaults" {
  description = "Map of default values which will be used for each item."
  type        = any
  default     = {}
}

variable "items" {
  description = "Maps of items to create a wrapper from. Values are passed through to the module."
  type        = any
  default     = {}
}
